package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/hex"
   "errors"
   "log/slog"
   "net/url"
   "os"
)

func (h HttpStream) key(key_id []byte) ([]byte, error) {
   private_key, err := os.ReadFile(h.PrivateKey)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(h.ClientId)
   if err != nil {
      return nil, err
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
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
   slog.Debug("CDM", "key", hex.EncodeToString(key))
   return key, nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type HttpStream struct {
   ClientId string
   PrivateKey string
   Name encoding.Namer
   base *url.URL
   Poster widevine.Poster
}
