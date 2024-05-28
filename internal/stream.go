package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "crypto/tls"
   "encoding/hex"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strings"
)

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
   var protect protection
   err = protect.init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
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

func (s Stream) Download(rep *dash.Representation) error {
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Representation.Ext")
   }
   base := rep.GetAdaptationSet().GetPeriod().GetMpd().BaseUrl.URL
   if v, ok := rep.GetSegmentTemplate(); ok {
      if v, ok := v.GetInitialization(rep); ok {
         return s.segment_template(rep, base, v, ext)
      }
   }
   return s.segment_base(rep.SegmentBase, base, *rep.BaseUrl, ext)
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
   var protect protection
   err = protect.init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := s.key(protect)
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

func (s Stream) key(protect protection) ([]byte, error) {
   if protect.key_id == nil {
      return nil, nil
   }
   private_key, err := os.ReadFile(s.PrivateKey)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(s.ClientId)
   if err != nil {
      return nil, err
   }
   if protect.pssh == nil {
      protect.pssh = widevine.PSSH(protect.key_id, nil)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, protect.pssh)
   if err != nil {
      return nil, err
   }
   key, err := module.Key(s.Poster, protect.key_id)
   if err != nil {
      return nil, err
   }
   slog.Debug("CDM", "key", hex.EncodeToString(key))
   return key, nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type Stream struct {
   ClientId string
   PrivateKey string
   Name text.Namer
   Poster widevine.Poster
}
