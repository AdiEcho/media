package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "errors"
   "net/url"
   "os"
)

func (h HttpStream) key(pssh []byte) ([]byte, error) {
   client_id, err := os.ReadFile(h.ClientId)
   if err != nil {
      return nil, err
   }
   private_key, err := os.ReadFile(h.PrivateKey)
   if err != nil {
      return nil, err
   }
   protect := widevine.PSSH{Data: pssh}
   if err := protect.Consume(); err != nil {
      return nil, err
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      return nil, err
   }
   license, err := module.License(h.Poster)
   if err != nil {
      return nil, err
   }
   key, ok := module.Key(license)
   if !ok {
      return nil, errors.New("widevine.CDM.Key")
   }
   return key, nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type HttpStream struct {
   ClientId string
   Name encoding.Namer
   Poster widevine.Poster
   PrivateKey string
   base *url.URL
}
