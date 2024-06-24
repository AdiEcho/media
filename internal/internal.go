package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/text"
   "crypto/tls"
   "errors"
   "io"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strings"
)

func DASH(req *http.Request) ([]*dash.Representation, error) {
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
   var media dash.Mpd
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   err = media.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   if media.BaseUrl == nil {
      media.BaseUrl = &dash.URL{res.Request.URL}
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

func (s Stream) segment_template(
   rep *dash.Representation,
   base *url.URL,
   initial string,
   ext string,
) error {
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
      s, err := text.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(text.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   err = s.init_protect(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   template, ok := rep.GetSegmentTemplate()
   if !ok {
      return errors.New("GetSegmentTemplate")
   }
   media, err := template.GetMedia(rep)
   if err != nil {
      return err
   }
   client := http.Client{ // github.com/golang/go/issues/18639
      Transport: &http.Transport{
         Proxy: http.ProxyFromEnvironment,
         TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
      },
   }
   var meter text.ProgressMeter
   meter.Set(len(media))
   var log text.LogLevel
   log.SetTransport(false)
   defer log.SetTransport(true)
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
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
func (s Stream) segment_base(
   segment *dash.SegmentBase,
   base *url.URL,
   initial string,
   ext string,
) error {
   req, err := http.NewRequest("", initial, nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
   data, _ := segment.Initialization.Range.MarshalText()
   req.Header.Set("Range", "bytes=" + string(data))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := text.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(text.Clean(s) + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   err = s.init_protect(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   references, err := write_sidx(req, segment.IndexRange)
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(references))
   var log text.LogLevel
   log.SetTransport(false)
   defer log.SetTransport(true)
   for _, reference := range references {
      segment.IndexRange.Start = segment.IndexRange.End + 1
      segment.IndexRange.End += uint64(reference.ReferencedSize())
      data, _ := segment.IndexRange.MarshalText()
      err := func() error {
         req.Header.Set("Range", "bytes=" + string(data))
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

func (s Stream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := text.Name(s.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(text.Clean(s) + ".vtt")
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

var Forward = ForwardedFor{
{"Argentina", "186.128.0.0"},
{"Australia", "1.128.0.0"},
{"Bolivia", "179.58.0.0"},
{"Brazil", "179.192.0.0"},
{"Canada", "99.224.0.0"},
{"Chile", "191.112.0.0"},
{"Colombia", "181.128.0.0"},
{"Costa Rica", "201.192.0.0"},
{"Denmark", "2.104.0.0"},
{"Ecuador", "186.68.0.0"},
{"Egypt", "197.32.0.0"},
{"Germany", "53.0.0.0"},
{"Guatemala", "190.56.0.0"},
{"India", "106.192.0.0"},
{"Indonesia", "39.192.0.0"},
{"Ireland", "87.32.0.0"},
{"Italy", "79.0.0.0"},
{"Latvia", "78.84.0.0"},
{"Malaysia", "175.136.0.0"},
{"Mexico", "189.128.0.0"},
{"Netherlands", "145.160.0.0"},
{"New Zealand", "49.224.0.0"},
{"Norway", "88.88.0.0"},
{"Peru", "190.232.0.0"},
{"Russia", "95.24.0.0"},
{"South Africa", "105.0.0.0"},
{"South Korea", "175.192.0.0"},
{"Spain", "88.0.0.0"},
{"Sweden", "78.64.0.0"},
{"Taiwan", "120.96.0.0"},
{"United Kingdom", "25.0.0.0"},
{"Venezuela", "190.72.0.0"},
}
