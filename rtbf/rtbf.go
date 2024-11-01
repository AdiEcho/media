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

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   a.Path = strings.TrimPrefix(s, "auvio.rtbf.be")
   return nil
}

type AuvioAuth struct {
   SessionToken string
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
   u.Scheme = "https"
   u.Host = "rbm-rtbf.live.ott.irdeto.com"
   u.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
   u.RawQuery = url.Values{
      "contentId":  {e.AssetId},
      "ls_session": {e.PlayToken},
   }.Encode()
   return u.String(), true
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

type Entitlement struct {
   AssetId   string
   PlayToken string
   Formats   []struct {
      Format       string
      MediaLocator string
   }
}

type Namer struct {
   Page *AuvioPage
}

// its just not available from what I can tell
func (Namer) Year() int {
   return 0
}

func (n Namer) Episode() int {
   return n.Page.Subtitle.Episode
}

func (n Namer) Season() int {
   return n.Page.Title.Season
}

func (n Namer) Show() string {
   if v := n.Page.Title; v.Season >= 1 {
      return v.Title
   }
   return ""
}

func (n Namer) Title() string {
   if v := n.Page.Subtitle; v.Episode >= 1 {
      return v.Subtitle
   }
   return n.Page.Title.Title
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

type WebToken struct {
   ErrorMessage string
   IdToken      string `json:"id_token"`
}

///

func (a *AuvioLogin) New(id, password string) ([]byte, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.login", url.Values{
         "APIKey":   {api_key},
         "loginID":  {id},
         "password": {password},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if a != nil {
      err = a.Unmarshal(data)
      if err != nil {
         return nil, err
      }
   }
   return data, nil
}

func (a Address) Page() (*AuvioPage, error) {
   resp, err := http.Get(
      "https://bff-service.rtbf.be/auvio/v1.23/pages" + a.Path,
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      Data struct {
         Content AuvioPage
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.Content, nil
}

func (w *WebToken) Auth() (*AuvioAuth, error) {
   var value struct {
      Device struct {
         DeviceId string `json:"deviceId"`
         Type     string `json:"type"`
      } `json:"device"`
      Jwt string `json:"jwt"`
   }
   value.Device.Type = "WEB"
   value.Jwt = w.IdToken
   body, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   web := func() string {
      var b strings.Builder
      b.WriteString("https://exposure.api.redbee.live")
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin")
      return b.String()
   }
   resp, err := http.Post(
      web(), "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   auth := &AuvioAuth{}
   err = json.NewDecoder(resp.Body).Decode(auth)
   if err != nil {
      return nil, err
   }
   return auth, nil
}

type Address struct {
   Path string
}

func (a Address) String() string {
   return a.Path
}

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

type AuvioPage struct {
   AssetId  string
   Media *struct {
      AssetId string
   }
   Subtitle Subtitle
   Title    Title
}

func (a *AuvioPage) GetAssetId() (string, bool) {
   if v := a.AssetId; v != "" {
      return v, true
   }
   if v := a.Media; v != nil {
      return v.AssetId, true
   }
   return "", false
}

func (a *AuvioAuth) Entitlement(asset_id string) (*Entitlement, error) {
   req, err := http.NewRequest("", "https://exposure.api.redbee.live", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement/")
      b.WriteString(asset_id)
      b.WriteString("/play")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization":   {"Bearer " + a.SessionToken},
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

func (a *AuvioLogin) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, a)
   if err != nil {
      return err
   }
   if v := a.ErrorMessage; v != "" {
      return errors.New(v)
   }
   return nil
}

type AuvioLogin struct {
   ErrorMessage string
   SessionInfo  struct {
      CookieValue string
   }
}

func (a *AuvioLogin) Token() (*WebToken, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.getJWT", url.Values{
         "APIKey": {api_key},
         "login_token": {a.SessionInfo.CookieValue},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var web WebToken
   err = json.NewDecoder(resp.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if v := web.ErrorMessage; v != "" {
      return nil, errors.New(v)
   }
   return &web, nil
}
