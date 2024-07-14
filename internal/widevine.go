package internal

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "log/slog"
   "os"
)

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
      s.pssh = widevine.Pssh{KeyId: s.key_id}.Encode()
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
      "CDM", "ID", hex.EncodeToString(s.key_id), "key", hex.EncodeToString(key),
   )
   return key, nil
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

func (s Stream) file(ext string) (*os.File, error) {
   name, err := text.Name(s.Name)
   if err != nil {
      return nil, err
   }
   return os.Create(text.Clean(name) + ext)
}
