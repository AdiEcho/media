package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "crypto/tls"
   "errors"
   "io"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strconv"
   "strings"
)

func (s Stream) segment_template(
   ext, initial string, rep *dash.Representation,
) error {
   base, err := url.Parse(rep.GetAdaptationSet().GetPeriod().GetMpd().BaseURL)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("GET", initial, nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   var protect protection
   err = protect.init(res.Body, file)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   template, ok := rep.GetSegmentTemplate()
   if !ok {
      return errors.New("GetSegmentTemplate")
   }
   media, err := template.GetMedia(rep)
   if err != nil {
      return err
   }
   meter.Set(len(media))
   client := http.Client{ // github.com/golang/go/issues/18639
      Transport: &http.Transport{
         Proxy: http.ProxyFromEnvironment,
         TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
      },
   }
   for _, medium := range media {
      req.URL, err = base.Parse(medium)
      if err != nil {
         return err
      }
      err := func() error {
         res, err := client.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            var b strings.Builder
            res.Write(&b)
            return errors.New(b.String())
         }
         return write_segment(meter.Reader(res), file, key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s *Stream) DASH(req *http.Request) ([]*dash.Representation, error) {
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   switch res.Status {
   case "200 OK", "403 OK":
   default:
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var media dash.MPD
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   err = media.Unmarshal(text)
   if err != nil {
      return nil, err
   }
   if media.BaseURL == "" {
      media.BaseURL = res.Request.URL.String()
   }
   var reps []*dash.Representation
   for _, v := range media.Period {
      seconds, err := v.Seconds()
      if err != nil {
         return nil, err
      }
      for _, v := range v.AdaptationSet {
         for _, v := range v.Representation {
            if seconds > 9 {
               if _, ok := v.Ext(); ok {
                  reps = append(reps, v)
               }
            }
         }
      }
   }
   slices.SortFunc(reps, func(a, b *dash.Representation) int {
      return int(a.Bandwidth - b.Bandwidth)
   })
   return reps, nil
}

func (s Stream) Download(rep *dash.Representation) error {
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Ext")
   }
   if v, ok := rep.GetSegmentTemplate(); ok {
      if v, ok := v.GetInitialization(rep); ok {
         return s.segment_template(ext, v, rep)
      }
   }
   return s.segment_base(ext, *rep.BaseURL, rep)
}


func (s Stream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ".vtt")
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   _, err = file.ReadFrom(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (s Stream) segment_base(
   ext, base_url string, rep *dash.Representation,
) error {
   sb := rep.SegmentBase
   req, err := http.NewRequest("GET", base_url, nil)
   if err != nil {
      return err
   }
   req.Header.Set("Range", "bytes=" + string(sb.Initialization.Range))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := encoding.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(encoding.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   var protect protection
   err = protect.init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
   if err != nil {
      return err
   }
   references, err := write_sidx(base_url, sb.IndexRange)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   meter.Set(len(references))
   var start uint64
   end, err := func() (uint64, error) {
      _, s, _ := sb.IndexRange.Cut()
      return strconv.ParseUint(s, 10, 64)
   }()
   if err != nil {
      return err
   }
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, reference := range references {
      start = end + 1
      end += uint64(reference.ReferencedSize())
      bytes := func() string {
         b := []byte("bytes=")
         b = strconv.AppendUint(b, start, 10)
         b = append(b, '-')
         b = strconv.AppendUint(b, end, 10)
         return string(b)
      }()
      err := func() error {
         req, err := http.NewRequest("GET", base_url, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", bytes)
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusPartialContent {
            return errors.New(res.Status)
         }
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

var Forward = ForwardedFor{
   {"Argentina", "181.0.0.0"},
   {"Australia", "1.128.0.0"},
   {"Bolivia", "179.58.0.0"},
   {"Brazil", "189.0.0.0"},
   {"Canada", "99.224.0.0"},
   {"Chile", "191.112.0.0"},
   {"Colombia", "181.128.0.0"},
   {"Costa Rica", "201.192.0.0"},
   {"Denmark", "87.48.0.0"},
   {"Ecuador", "186.68.0.0"},
   {"Germany", "53.0.0.0"},
   {"Guatemala", "190.148.0.0"},
   {"Ireland", "87.32.0.0"},
   {"Italy", "79.0.0.0"},
   {"Mexico", "189.128.0.0"},
   {"Norway", "88.88.0.0"},
   {"Peru", "190.232.0.0"},
   {"South Africa", "41.0.0.0"},
   {"Spain", "88.0.0.0"},
   {"Sweden", "78.64.0.0"},
   {"United Kingdom", "25.0.0.0"},
   {"Venezuela", "186.88.0.0"},
}

type ForwardedFor []struct {
   Country string
   IP string
}

func (f ForwardedFor) String() string {
   var b strings.Builder
   for _, each := range f {
      if b.Len() >= 1 {
         b.WriteByte('\n')
      }
      b.WriteString(each.Country)
      b.WriteByte(' ')
      b.WriteString(each.IP)
   }
   return b.String()
}

