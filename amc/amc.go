package amc

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
   "time"
)

type CurrentVideo struct {
   Meta struct {
      Airdate time.Time // 1996-01-01T00:00:00.000Z
      EpisodeNumber int
      Season int `json:",string"`
      ShowTitle string
   }
   Text struct {
      Title string
   }
}

func (c *CurrentVideo) Episode() int {
   return c.Meta.EpisodeNumber
}

func (c *CurrentVideo) Show() string {
   return c.Meta.ShowTitle
}

func (c CurrentVideo) Season() int {
   return c.Meta.Season
}

func (c CurrentVideo) Title() string {
   return c.Text.Title
}

func (c CurrentVideo) Year() int {
   return c.Meta.Airdate.Year()
}

type DataSource struct {
   KeySystems *struct {
      Widevine struct {
         LicenseUrl string `json:"license_url"`
      } `json:"com.widevine.alpha"`
   } `json:"key_systems"`
   Src string
   Type string
}

func (Playback) DashError() error {
   return errors.New("Playback.DashError")
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (a *Authorization) Playback(nid string) (*Playback, error) {
   var req_body struct {
      AdTags struct {
         Lat int `json:"lat"`
         Mode string `json:"mode"`
         Ppid int `json:"ppid"`
         PlayerHeight int `json:"playerHeight"`
         PlayerWidth int `json:"playerWidth"`
         Url string `json:"url"`
      } `json:"adtags"`
   }
   req_body.AdTags.Mode = "on-demand"
   req_body.AdTags.Url = "-"
   data, err := json.Marshal(req_body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/playback-id/api/v1/playback/" + nid
   req.Header = http.Header{
      "authorization": {"Bearer " + a.AccessToken},
      "content-type": {"application/json"},
      "x-amcn-device-ad-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-service-id": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-ccpa-do-not-sell": {"doNotPassData"},
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
   var play Playback
   err = json.NewDecoder(resp.Body).Decode(&play)
   if err != nil {
      return nil, err
   }
   play.AmcnBcJwt = resp.Header.Get("x-amcn-bc-jwt")
   return &play, nil
}

func (p *Playback) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("bcov-auth", p.AmcnBcJwt)
   return head, nil
}

func (p *Playback) Dash() (*DataSource, bool) {
   for _, source := range p.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source, true
      }
   }
   return nil, false
}

type Playback struct {
   AmcnBcJwt string `json:"-"`
   Data struct {
      PlaybackJsonData struct {
         Sources []DataSource
      }
   }
}

func (p *Playback) RequestUrl() (string, bool) {
   if v, ok := p.Dash(); ok {
      return v.KeySystems.Widevine.LicenseUrl, true
   }
   return "", false
}

func cache_hash() string {
   return base64.StdEncoding.EncodeToString([]byte("ff="))
}

type Address struct {
   Nid string
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (a *Address) Set(text string) error {
   var found bool
   _, a.Path, found = strings.Cut(text, "amcplus.com")
   if !found {
      return errors.New("amcplus.com")
   }
   _, a.Nid, found = strings.Cut(a.Path, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}

func (a *Authorization) Unmarshal() error {
   var value struct {
      Data Authorization
   }
   err := json.Unmarshal(a.Raw, &value)
   if err != nil {
      return err
   }
   *a = value.Data
   return nil
}

func (a *Authorization) Login(email, password string) error {
   body, err := json.Marshal(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/login"
   req.Header = http.Header{
      "authorization": {"Bearer " + a.AccessToken},
      "content-type": {"application/json"},
      "x-amcn-device-ad-id": {"-"},
      "x-amcn-device-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-service-group-id": {"10"},
      "x-amcn-service-id": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-ccpa-do-not-sell": {"doNotPassData"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *Authorization) Refresh() error {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/refresh"
   req.Header.Set("authorization", "Bearer " + a.RefreshToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *Authorization) Unauth() error {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/unauth"
   req.Header = http.Header{
      "x-amcn-device-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-tenant": {"amcn"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

type Authorization struct {
   AccessToken string `json:"access_token"`
   RefreshToken string `json:"refresh_token"`
   Raw []byte `json:"-"`
}

func (ContentCompiler) VideoError() error {
   return errors.New("ContentCompiler.Video")
}

///

func (a *Authorization) Content(path string) (*ContentCompiler, error) {
   req, err := http.NewRequest("", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "authorization": {"Bearer " + a.AccessToken},
      "x-amcn-cache-hash": {cache_hash()},
      "x-amcn-network": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-amcn-user-cache-hash": {cache_hash()},
   }
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/content-compiler-cr/api/v1/content/amcn/amcplus/path")
      if strings.HasPrefix(path, "/movies/") {
         b.WriteString("/watch")
      }
      b.WriteString(path)
      return b.String()
   }()
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
   content := &ContentCompiler{}
   err = json.NewDecoder(resp.Body).Decode(content)
   if err != nil {
      return nil, err
   }
   return content, nil
}

func (c *ContentCompiler) Video() (*CurrentVideo, bool) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         return &child.Properties.CurrentVideo, true
      }
   }
   return nil, false
}

type ContentCompiler struct {
   Data   struct {
      Children []struct {
         Properties struct {
            CurrentVideo CurrentVideo
         }
         Type string
      }
   }
}
