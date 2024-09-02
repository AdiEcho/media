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

func (a *AuvioAuth) Entitlement(page *AuvioPage) (*Entitlement, error) {
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

type AuvioAuth struct {
   SessionToken string
}

func (a *AuvioPage) New(path string) error {
   resp, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   var value struct {
      Data struct {
         Content AuvioPage
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   *a = value.Data.Content
   return nil
}

func (a *AuvioPage) asset_id() string {
   if a.AssetId != "" {
      return a.AssetId
   }
   return a.Media.AssetId
}

type AuvioPage struct {
   AssetId  string
   Media struct {
      AssetId string
   }
   Subtitle Subtitle
   Title    Title
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

// its just not available from what I can tell
func (Namer) Year() int {
   return 0
}

type Namer struct {
   Page AuvioPage
}

func (n *Namer) Episode() int {
   return n.Page.Subtitle.Episode
}

func (n *Namer) Season() int {
   return n.Page.Title.Season
}

func (n *Namer) Show() string {
   if v := n.Page.Title; v.Season >= 1 {
      return v.Title
   }
   return ""
}

func (n *Namer) Title() string {
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
   auth := &AuvioAuth{}
   err = json.NewDecoder(resp.Body).Decode(auth)
   if err != nil {
      return nil, err
   }
   return auth, nil
}
