package rtbf

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (a *AuvioPage) Episode() int {
   return a.Content.Subtitle.Episode
}

// its just not available from what I can tell
func (AuvioPage) Year() int {
   return 0
}

func (a *AuvioPage) Show() string {
   if v := a.Content.Title; v.Season >= 1 {
      return v.Title
   }
   return ""
}

func (a *AuvioPage) Title() string {
   if v := a.Content.Subtitle; v.Episode >= 1 {
      return v.Subtitle
   }
   return a.Content.Title.Title
}

func (a *AuvioPage) asset_id() string {
   if v := a.Content.AssetId; v != "" {
      return v
   }
   return a.Content.Media.AssetId
}

func (a *AuvioPage) Season() int {
   return a.Content.Title.Season
}

type AuvioPage struct {
   Content struct {
      AssetId  string
      Media struct {
         AssetId string
      }
      Subtitle Subtitle
      Title    Title
   }
}

type Subtitle struct {
   Episode  int
   Subtitle string
}

// json.data.content.subtitle = "06 - Les ombres de la guerre";
// json.data.content.subtitle = "Avec Rosamund Pike";
func (s *Subtitle) UnmarshalText(text []byte) error {
   s.Subtitle = string(text)
   if before, after, ok := strings.Cut(s.Subtitle, " - "); ok {
      if episode, err := strconv.Atoi(before); err == nil {
         s.Episode = episode
         s.Subtitle = after
      }
   }
   return nil
}

type Title struct {
   Season int
   Title  string
}

// json.data.content.title = "Grantchester S01";
// json.data.content.title = "I care a lot";
func (t *Title) UnmarshalText(text []byte) error {
   t.Title = string(text)
   if before, after, ok := strings.Cut(t.Title, " S"); ok {
      if season, err := strconv.Atoi(after); err == nil {
         t.Title = before
         t.Season = season
      }
   }
   return nil
}

type Entitlement struct {
   AssetId   string
   PlayToken string
   Formats   []struct {
      Format       string
      MediaLocator string
   }
}

func (Entitlement) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Entitlement) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (Entitlement) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("content-type", "application/x-protobuf")
   return head, nil
}

func (e *Entitlement) Dash() (string, bool) {
   for _, format := range e.Formats {
      if format.Format == "DASH" {
         return format.MediaLocator, true
      }
   }
   return "", false
}

func (e *Entitlement) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "rbm-rtbf.live.ott.irdeto.com"
   u.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
   u.Scheme = "https"
   u.RawQuery = url.Values{
      "contentId":  {e.AssetId},
      "ls_session": {e.PlayToken},
   }.Encode()
   return u.String(), true
}

type AuvioGigya struct {
   SessionToken string
}

type AuvioToken struct {
   ErrorMessage string
   IdToken      string `json:"id_token"`
}

func (g *AuvioGigya) Entitlement(page *AuvioPage) (*Entitlement, error) {
   req, err := http.NewRequest("", "https://exposure.api.redbee.live", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement/")
      b.WriteString(page.asset_id())
      b.WriteString("/play")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization":   {"Bearer " + g.SessionToken},
      "x-forwarded-for": {"91.90.123.17"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   title := &Entitlement{}
   err = json.NewDecoder(resp.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

func NewPage(path string) (*AuvioPage, error) {
   resp, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var data struct {
      Data AuvioPage
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &data.Data, nil
}

func (t *AuvioLogin) New(id, password string) error {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.login", url.Values{
         "APIKey":   {api_key},
         "loginID":  {id},
         "password": {password},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   t.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

type AuvioLogin struct {
   CookieValue string
   Raw []byte
}

func (t *AuvioLogin) Unmarshal() error {
   var data struct {
      ErrorMessage string
      SessionInfo  struct {
         CookieValue string
      }
   }
   err := json.Unmarshal(t.Raw, &data)
   if err != nil {
      return err
   }
   if v := data.ErrorMessage; v != "" {
      return errors.New(v)
   }
   t.CookieValue = data.SessionInfo.CookieValue
   return nil
}

func (t *AuvioLogin) AuvioToken() (*AuvioToken, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.getJWT", url.Values{
         "APIKey": {api_key},
         "login_token": {t.CookieValue},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var web AuvioToken
   err = json.NewDecoder(resp.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if v := web.ErrorMessage; v != "" {
      return nil, errors.New(v)
   }
   return &web, nil
}

func (w *AuvioToken) Login() (*AuvioGigya, error) {
   var data struct {
      Device struct {
         DeviceId string `json:"deviceId"`
         Type     string `json:"type"`
      } `json:"device"`
      Jwt string `json:"jwt"`
   }
   data.Device.Type = "WEB"
   data.Jwt = w.IdToken
   body, err := json.Marshal(data)
   if err != nil {
      return nil, err
   }
   address := func() string {
      var b strings.Builder
      b.WriteString("https://exposure.api.redbee.live")
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin")
      return b.String()
   }
   resp, err := http.Post(
      address(), "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   login := &AuvioGigya{}
   err = json.NewDecoder(resp.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}
