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
   "strings"
)

func DASH(req *http.Request) (chan dash.Period, error) {
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   switch resp.Status {
   case "200 OK", "403 OK":
   default:
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var media dash.Mpd
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   err = media.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   if media.BaseUrl == nil {
      media.BaseUrl = &dash.Url{resp.Request.URL}
   }
   return media.GetPeriod(), nil
}

func (s Stream) segment_base(
   ext string,
   initial, base *url.URL,
   segment *dash.SegmentBase,
) error {
   req := new(http.Request)
   req.URL = base.ResolveReference(initial)
   data, _ := segment.Initialization.Range.MarshalText()
   req.Header.Set("Range", "bytes=" + string(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
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
   err = s.init_protect(file, resp.Body)
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
         resp, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusPartialContent {
            return errors.New(resp.Status)
         }
         return write_segment(file, meter.Reader(resp), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s Stream) segment_template(
   ext, initial string,
   base *url.URL,
   rep dash.Representation,
) error {
   req, err := http.NewRequest("", initial, nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
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
   err = s.init_protect(file, resp.Body)
   if err != nil {
      return err
   }
   key, err := s.key()
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
   var log text.LogLevel
   log.SetTransport(false)
   defer log.SetTransport(true)
   media := rep.Media()
   meter.Set(len(media))
   for _, medium := range media {
      req.URL, err = base.Parse(medium)
      if err != nil {
         return err
      }
      err := func() error {
         resp, err := client.Do(req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusOK {
            var b strings.Builder
            resp.Write(&b)
            return errors.New(b.String())
         }
         return write_segment(file, meter.Reader(resp), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s Stream) TimedText(url string) error {
   resp, err := http.Get(url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
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
   _, err = file.ReadFrom(resp.Body)
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
