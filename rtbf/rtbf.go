package rtbf

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type embed_media struct {
   Data struct {
      AssetId string
      Program *struct {
         Title string
      }
      Subtitle string
      Title string
   }
   Meta struct {
      SmartAds struct {
         CTE number
         CTS number
      }
   }
}

func (e *embed_media) New(media int64) error {
   address := func() string {
      b := []byte("https://bff-service.rtbf.be/auvio/v1.23/embed/media/")
      b = strconv.AppendInt(b, media, 10)
      b = append(b, "?userAgent"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(e)
}

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (a *accounts_login) New(id, password string) error {
   body := url.Values{
      "APIKey": {api_key},
      "loginID": {id},
      "password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.login",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   err = json.NewDecoder(res.Body).Decode(a)
   if err != nil {
      return err
   }
   if v := a.ErrorMessage; v != "" {
      return errors.New(v)
   }
   return nil
}

type accounts_login struct {
   ErrorMessage string
   SessionInfo struct {
      CookieValue string
   }
}

func (a *accounts_login) unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

func (a accounts_login) marshal() ([]byte, error) {
   return json.Marshal(a)
}

func (a accounts_login) token() (*web_token, error) {
   body := url.Values{
      "APIKey": {api_key},
      // from /accounts.login
      "login_token": {a.SessionInfo.CookieValue},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.getJWT",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var web web_token
   err = json.NewDecoder(res.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if v := web.ErrorMessage; v != "" {
      return nil, errors.New(v)
   }
   return &web, nil
}

func (e embed_media) Episode() int {
   return int(e.Meta.SmartAds.CTE)
}

func (e embed_media) Season() int {
   return int(e.Meta.SmartAds.CTS)
}

func (e embed_media) Show() string {
   if v := e.Data.Program; v != nil {
      return v.Title
   }
   return ""
}

func (e embed_media) Title() string {
   if e.Data.Program != nil {
      // json.data.subtitle = "06 - Les ombres de la guerre";
      _, after, _ := strings.Cut(e.Data.Subtitle, " - ")
      return after
   }
   // json.data.title = "I care a lot";
   return e.Data.Title
}

// its just not available from what I can tell
func (embed_media) Year() int {
   return 0
}

type gigya_login struct {
   SessionToken string
}

type number int

func (n *number) UnmarshalText(text []byte) error {
   if len(text) >= 1 {
      i, err := strconv.Atoi(string(text))
      if err != nil {
         return err
      }
      *n = number(i)
   }
   return nil
}

type web_token struct {
   ErrorMessage string
   IdToken string `json:"id_token"`
}

func (w web_token) login() (*gigya_login, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Device struct {
            DeviceId string `json:"deviceId"`
            Type string `json:"type"`
         } `json:"device"`
         JWT string `json:"jwt"`
      }
      s.Device.Type = "WEB"
      s.JWT = w.IdToken
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://exposure.api.redbee.live", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin"
   req.Header.Set("content-type", "application/json")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(gigya_login)
   err = json.NewDecoder(res.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}
