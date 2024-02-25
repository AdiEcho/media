package peacock

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (auth_tokens) id_session() string {
   return "E0-AAAADBDXsYWL1va0ydRMpsRHQz8eda1M0SzgUf3E9nz+KwOkmuki3oNthDYAayhFQL3CtFJcpPVB7HcM0GWA8StjtObMQvR3HWx016ugrvQ6IovMVFV/1Tf8VrbUlVmxYoCi"
}

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
            // request will work without this, but then `/video/playouts/vod`
            // will fail with
            // {"errorCode":"OVP_00311","description":"Unknown deviceId"}
            ID string `json:"id"`
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
      /*
      pass
      s.Device.ID = "AAAAAAAAAAAAAAAAAAAA"
      peacocktv.com/help/article/why-am-i-seeing-an-error-that-i-ve-reached-the-simultaneous-streams-limit
      {"errorCode":"OVP_00014",
      "description":"Maximum number of streaming devices exceeded"}
      Date: Sun, 25 Feb 2024 05:34:27 GMT
      1 minute 1135p fail
      2 minute 1136p fail
      4 minute 1138p fail
      8 minute 1142p fail
      16 minute 1150p fail
      32 minute 1206a
      */
      s.Device.ID = "AAAAAAAAAAAAAAAAAAAB"
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://play.ovp.peacocktv.com/auth/tokens",
      bytes.NewReader(body),
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

type auth_tokens struct {
   UserToken string
}
