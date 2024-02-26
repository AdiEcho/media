package peacock

import (
   "bytes"
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha1"
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "slices"
   "strings"
   "time"
)

const (
   app_id = "NBCU-ANDROID-v3"
   signature_key = "JuLQgyFz9n89D9pxcN6ZWZXKWfgj2PNBUb32zybj"
   sig_version = "1.0"
)

func (s *sign_in) unmarshal(b []byte) error {
   return json.Unmarshal(b, &s.cookie)
}

func (s sign_in) marshal() ([]byte, error) {
   return json.Marshal(s.cookie)
}

type sign_in struct {
   cookie *http.Cookie
}

// userToken is good for one day
type auth_tokens struct {
   UserToken string
}

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
      "POST", "https://ovp.peacocktv.com/auth/tokens", bytes.NewReader(body),
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

func sign(method, path string, head http.Header, body []byte) string {
   timestamp := time.Now().Unix()
   text_headers := func() string {
      var s []string
      for k := range head {
         k = strings.ToLower(k)
         if strings.HasPrefix(k, "x-skyott-") {
            s = append(s, k + ": " + head.Get(k) + "\n")
         }
      }
      slices.Sort(s)
      return strings.Join(s, "")
   }()
   headers_md5 := md5.Sum([]byte(text_headers))
   payload_md5 := md5.Sum(body)
   signature := func() string {
      h := hmac.New(sha1.New, []byte(signature_key))
      fmt.Fprintln(h, method)
      fmt.Fprintln(h, path)
      fmt.Fprintln(h)
      fmt.Fprintln(h, app_id)
      fmt.Fprintln(h, sig_version)
      fmt.Fprintf(h, "%x\n", headers_md5)
      fmt.Fprintln(h, timestamp)
      fmt.Fprintf(h, "%x\n", payload_md5)
      hashed := h.Sum(nil)
      return base64.StdEncoding.EncodeToString(hashed[:])
   }
   sky_ott := func() string {
      b := []byte("SkyOTT")
      // must be quoted
      b = fmt.Appendf(b, " client=%q", app_id)
      // must be quoted
      b = fmt.Appendf(b, ",signature=%q", signature())
      // must be quoted
      b = fmt.Appendf(b, `,timestamp="%v"`, timestamp)
      // must be quoted
      b = fmt.Appendf(b, ",version=%q", sig_version)
      return string(b)
   }
   return sky_ott()
}
