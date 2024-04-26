package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/hex"
   "errors"
   "log/slog"
   "os"
)

func (s Stream) key(key_id []byte) ([]byte, error) {
   if key_id == nil {
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
   var module widevine.CDM
   err = module.New(private_key, client_id, key_id)
   if err != nil {
      return nil, err
   }
   license, err := module.License(s.Poster)
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
type Stream struct {
   ClientId string
   PrivateKey string
   Name encoding.Namer
   Poster widevine.Poster
}
