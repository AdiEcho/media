package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/text"
   "crypto/tls"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (s *Stream) segment_template(
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
      data, err := io.ReadAll(resp.Body)
      if err != nil {
         return err
      }
      data, err = s.init_protect(data)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
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
   client := http.Client{ // github.com/golang/go/issues/18639
      Transport: &http.Transport{
         Proxy: http.ProxyFromEnvironment,
         TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
      },
   }
   for _, medium := range media {
      req.URL, err = base.Url.Parse(medium)
      if err != nil {
         return err
      }
      data, err := func() ([]byte, error) {
         resp, err := client.Do(req)
         if err != nil {
            return nil, err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusOK {
            var b strings.Builder
            resp.Write(&b)
            return nil, errors.New(b.String())
         }
         return io.ReadAll(meter.Reader(resp))
      }()
      if err != nil {
         return err
      }
      data, err = write_segment(data, key)
      if err != nil {
         return err
      }
      _, err = file.Write(data)
      if err != nil {
         return err
      }
   }
   return nil
}
