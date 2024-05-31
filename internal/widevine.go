package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/hex"
   "errors"
   "log/slog"
   "os"
)

func (s *Stream) Download(rep *dash.Representation) error {
   if v, ok := rep.Widevine(); ok {
      var err error
      s.pssh, err = base64.StdEncoding.DecodeString(v)
      if err != nil {
         return err
      }
   }
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
      s.pssh = widevine.PSSH(s.key_id, nil)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, s.pssh)
   if err != nil {
      return nil, err
   }
   key, err := module.Key(s.Poster, s.key_id)
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
   pssh []byte
   key_id []byte
}
