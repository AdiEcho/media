package peacock

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (a *auth_tokens) New() error {
   body, err := func() ([]byte, error) {
      var s struct {
         Auth struct {
            AuthScheme string `json:"authScheme"`
            Proposition string `json:"proposition"`
            Provider string `json:"provider"`
            ProviderTerritory string `json:"providerTerritory"`
         } `json:"auth"`
         Device struct {
            Platform string `json:"platform"`
            Type string `json:"type"`
         } `json:"device"`
      }
      s.Auth.AuthScheme = "MESSO"
      s.Auth.Proposition = "NBCUOTT"
      s.Auth.Provider = "NBCU"
      s.Auth.ProviderTerritory = "US"
      s.Device.Platform = "PC"
      s.Device.Type = "COMPUTER"
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://ovp.peacocktv.com/auth/tokens", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "content-type": {"application/vnd.tokens.v1+json"},
      "cookie": {"idsession=" + a.id_session()},
      "x-sky-signature": {sign(req.Method, req.URL.Path, nil, body)},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(res.Body).Decode(a)
}

func (auth_tokens) id_session() string {
   return "E0-AAAADBaA4I4bW1mAUJgD8HiYL3/d97POed3b0DcxA/VMs87+3JJYZA6V23xO2DTE9hgF1p2EUh6C5RWt0snSpc8+XZnCOetS3GsBlae8zfESQbonJTONmZxa4aypEqgYgWNW"
}

type auth_tokens struct {
   UserToken string
}
