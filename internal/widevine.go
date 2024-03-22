package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/widevine"
   "errors"
   "net/url"
   "os"
)

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type HttpStream struct {
   ClientId string
   Name encoding.Namer
   Poster widevine.Poster
   PrivateKey string
   base *url.URL
}

func (h HttpStream) key(rep dash.Representation) ([]byte, error) {
   client_id, err := os.ReadFile(h.ClientId)
   if err != nil {
      return nil, err
   }
   private_key, err := os.ReadFile(h.PrivateKey)
   if err != nil {
      return nil, err
   }
   var protect widevine.PSSH
   err = func() error {
      if v, ok := rep.PSSH(); ok {
         b, err := v.Decode()
         if err != nil {
            return err
         }
         return protect.New(b)
      }
      if v, ok := rep.Default_KID(); ok {
         protect.Key_ID, err = v.Decode()
         if err != nil {
            return err
         }
      }
      return nil
   }()
   if err != nil {
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
