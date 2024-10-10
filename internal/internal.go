package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "os"
   "strings"
)

func (s Stream) segment_template(
   ext, initial string, base *dash.BaseUrl, media []string,
) error {
   file, err := s.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("", initial, nil)
   if err != nil {
      return err
   }
   if initial != "" {
      req.URL = base.Url.ResolveReference(req.URL)
      resp, err := http.DefaultClient.Do(req)
      if err != nil {
         return err
      }
      defer resp.Body.Close()
      if resp.StatusCode != http.StatusOK {
         return errors.New(resp.Status)
      }
      buf, err := io.ReadAll(resp.Body)
      if err != nil {
         return err
      }
      buf, err = s.init_protect(buf)
      if err != nil {
         return err
      }
      _, err = file.Write(buf)
      if err != nil {
         return err
      }
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(media))
   var transport text.Transport
   transport.Set(false)
   defer transport.Set(true)
   for _, medium := range media {
      req.URL, err = base.Url.Parse(medium)
      if err != nil {
         return err
      }
      err := func() error {
         resp, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusOK {
            var b strings.Builder
            resp.Write(&b)
            return errors.New(b.String())
         }
         buf, err := io.ReadAll(meter.Reader(resp))
         if err != nil {
            return err
         }
         buf, err = write_segment(buf, key)
         if err != nil {
            return err
         }
         _, err = file.Write(buf)
         if err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s Stream) key() ([]byte, error) {
   if s.key_id == nil {
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
   if s.pssh == nil {
      s.pssh = widevine.Pssh{KeyId: s.key_id}.Marshal()
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, s.pssh)
   if err != nil {
      return nil, err
   }
   key, err := module.Key(s.Poster, s.key_id)
   if err != nil {
      return nil, err
   }
   slog.Info(
      "CDM",
      "PSSH", base64.StdEncoding.EncodeToString(s.pssh),
      "key", base64.StdEncoding.EncodeToString(key),
   )
   return key, nil
}

func (s Stream) Create(ext string) (*os.File, error) {
   name, err := text.Name(s.Name)
   if err != nil {
      return nil, err
   }
   return os.Create(text.Clean(name) + ext)
}

type ForwardedFor struct {
   Country string
   IP string
}

var Forward = []ForwardedFor{
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

func (s Stream) segment_base(
   ext string, base *dash.BaseUrl, segment *dash.SegmentBase,
) error {
   file, err := s.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   buf, _ := segment.Initialization.Range.MarshalText()
   var req http.Request
   req.URL = base.Url
   req.Header = http.Header{}
   // need to use Set for lower case
   req.Header.Set("range", "bytes=" + string(buf))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusPartialContent {
      return errors.New(resp.Status)
   }
   buf, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   buf, err = s.init_protect(buf)
   if err != nil {
      return err
   }
   _, err = file.Write(buf)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   references, err := write_sidx(&req, segment.IndexRange)
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(references))
   var transport text.Transport
   transport.Set(false)
   defer transport.Set(true)
   for _, reference := range references {
      segment.IndexRange.Start = segment.IndexRange.End + 1
      segment.IndexRange.End += uint64(reference.ReferencedSize())
      buf, _ := segment.IndexRange.MarshalText()
      err := func() error {
         req.Header.Set("range", "bytes=" + string(buf))
         resp, err := http.DefaultClient.Do(&req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusPartialContent {
            return errors.New(resp.Status)
         }
         buf, err := io.ReadAll(meter.Reader(resp))
         if err != nil {
            return err
         }
         buf, err = write_segment(buf, key)
         if err != nil {
            return err
         }
         _, err = file.Write(buf)
         if err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type Stream struct {
   ClientId string
   PrivateKey string
   Name text.Namer
   Poster widevine.Poster
   pssh []byte
   key_id []byte
}

func Dash(req *http.Request) ([]dash.Representation, error) {
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   buf, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return dash.Unmarshal(buf, resp.Request.URL)
}
