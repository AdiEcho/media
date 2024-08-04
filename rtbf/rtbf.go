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

func (a AuvioPage) Episode() int {
   return a.Content.Subtitle.Episode
}

func (a AuvioPage) Show() string {
   if v := a.Content.Title; v.Season >= 1 {
      return v.Title
   }
   return ""
}

func (a AuvioPage) Title() string {
   if v := a.Content.Subtitle; v.Episode >= 1 {
      return v.Subtitle
   }
   return a.Content.Title.Title
}

// its just not available from what I can tell
func (AuvioPage) Year() int {
   return 0
}

func (a AuvioPage) asset_id() string {
   if v := a.Content.AssetId; v != "" {
      return v
   }
   return a.Content.Media.AssetId
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

func (a AuvioPage) Season() int {
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

func (e Entitlement) Dash() (string, bool) {
   for _, format := range e.Formats {
      if format.Format == "DASH" {
         return format.MediaLocator, true
      }
   }
   return "", false
}

type Entitlement struct {
   AssetId   string
   PlayToken string
   Formats   []struct {
      Format       string
      MediaLocator string
   }
}

func (e Entitlement) RequestUrl() (string, bool) {
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

func (Entitlement) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("content-type", "application/x-protobuf")
   return head, nil
}

func (Entitlement) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Entitlement) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type GigyaLogin struct {
   SessionToken string
}

func (g GigyaLogin) Entitlement(page *AuvioPage) (*Entitlement, error) {
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

type WebToken struct {
   ErrorMessage string
   IdToken      string `json:"id_token"`
}

func (w WebToken) Login() (*GigyaLogin, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Device struct {
            DeviceId string `json:"deviceId"`
            Type     string `json:"type"`
         } `json:"device"`
         Jwt string `json:"jwt"`
      }
      s.Device.Type = "WEB"
      s.Jwt = w.IdToken
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   login := &GigyaLogin{}
   err = json.NewDecoder(resp.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}
