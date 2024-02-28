package amc

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func cache_hash() string {
   return base64.StdEncoding.EncodeToString([]byte("ff="))
}

func (a Authorization) Content(path string) (*ContentCompiler, error) {
   req, err := http.NewRequest("GET", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.s.Data.Access_Token},
      "X-Amcn-Cache-Hash": {cache_hash()},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Amcn-User-Cache-Hash": {cache_hash()},
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   content := new(ContentCompiler)
   if err := json.NewDecoder(res.Body).Decode(content); err != nil {
      return nil, err
   }
   return content, nil
}

type WebAddress struct {
   NID string
   Path string
}

func (w *WebAddress) Set(s string) error {
   var found bool
   _, w.Path, found = strings.Cut(s, "amcplus.com")
   if !found {
      return errors.New("amcplus.com")
   }
   _, w.NID, found = strings.Cut(w.Path, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}

func (w WebAddress) String() string {
   return w.Path
}

type Authorization struct {
   Raw []byte
   s struct {
      Data struct {
         Access_Token string
         Refresh_Token string
      }
   }
}

func (a Authorization) Playback(nid string) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         AdTags struct {
            Lat int `json:"lat"`
            Mode string `json:"mode"`
            PPID int `json:"ppid"`
            PlayerHeight int `json:"playerHeight"`
            PlayerWidth int `json:"playerWidth"`
            URL string `json:"url"`
         } `json:"adtags"`
      }
      s.AdTags.Mode = "on-demand"
      s.AdTags.URL = "-"
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/playback-id/api/v1/playback/" + nid
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.s.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-ID": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var play Playback
   play.header = res.Header
   if err := json.NewDecoder(res.Body).Decode(&play.body); err != nil {
      return nil, err
   }
   return &play, nil
}

type ContentCompiler struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (p Playback) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set("bcov-auth", p.header.Get("X-AMCN-BC-JWT"))
   return h, nil
}

type DataSource struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (Playback) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type Playback struct {
   header http.Header
   body struct {
      Data struct {
         PlaybackJsonData struct {
            Sources []DataSource
         }
      }
   }
}

func (p Playback) HttpsDash() (*DataSource, bool) {
   for _, s := range p.body.Data.PlaybackJsonData.Sources {
      if strings.HasPrefix(s.Src, "https://") {
         if s.Type == "application/dash+xml" {
            return &s, true
         }
      }
   }
   return nil, false
}

func (p Playback) RequestUrl() (string, bool) {
   if v, ok := p.HttpsDash(); ok {
      return v.Key_Systems.Widevine.License_URL, true
   }
   return "", false
}

func (a *Authorization) Unauth() error {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/unauth"
   req.Header = http.Header{
      "X-Amcn-Device-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   a.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *Authorization) Unmarshal() error {
   return json.Unmarshal(a.Raw, &a.s)
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
      "Authorization": {"Bearer " + a.s.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-ID": {"-"},
      "X-Amcn-Device-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Group-ID": {"10"},
      "X-Amcn-Service-ID": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   a.Raw, err = io.ReadAll(res.Body)
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
   req.Header.Set("Authorization", "Bearer " + a.s.Data.Refresh_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   a.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}
