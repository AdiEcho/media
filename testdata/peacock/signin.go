package peacock

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (s sign_in) auth() (*auth_tokens, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Auth struct {
            AuthScheme string `json:"authScheme"`
            Proposition string `json:"proposition"`
            Provider string `json:"provider"`
            ProviderTerritory string `json:"providerTerritory"`
         } `json:"auth"`
         Device struct {
            ID string `json:"id"`
            Platform string `json:"platform"`
            Type string `json:"type"`
         } `json:"device"`
      }
      s.Auth.AuthScheme = "MESSO"
      s.Auth.Proposition = "NBCUOTT"
      s.Auth.Provider = "NBCU"
      s.Auth.ProviderTerritory = "US"
      s.Device.Type = "COMPUTER"
      s.Device.Platform = "PC"
      // request will work without this, but then `/video/playouts/vod`
      // will fail with
      // {"errorCode":"OVP_00311","description":"Unknown deviceId"}
      // BE CAREFUL, changing this too often will result in a four hour block:
      // {"errorCode":"OVP_00014",
      // "description":"Maximum number of streaming devices exceeded"}
      s.Device.ID = "PC"
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://play.ovp.peacocktv.com/auth/tokens",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.cookie)
   req.Header.Set("content-type", "application/vnd.tokens.v1+json")
   req.Header.Set("x-sky-signature", sign(req.Method, req.URL.Path, nil, body))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   auth := new(auth_tokens)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

type sign_in struct {
   cookie *http.Cookie
}

func (s *sign_in) New(user, password string) error {
   body := url.Values{
      "userIdentifier": {user},
      "password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://rango.id.peacocktv.com/signin/service/international",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "X-Skyott-Proposition": {"NBCUOTT"},
      "X-Skyott-Provider": {"NBCU"},
      "X-Skyott-Territory": {"US"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   for _, cookie := range res.Cookies() {
      if cookie.Name == "idsession" {
         s.cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}
