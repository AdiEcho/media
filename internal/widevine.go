package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/hex"
   "log/slog"
   "os"
)

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
   Name encoding.Namer
   Poster widevine.Poster
}
