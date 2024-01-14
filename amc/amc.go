package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (r Raw_Auth) Unmarshal() (*Auth, error) {
   var a Auth
   err := json.Unmarshal(r, &a)
   if err != nil {
      return nil, err
   }
   return &a, nil
}

func Unauth() (Raw_Auth, error) {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
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
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

func (a Auth) Content(u URL) (*Content, error) {
   req, err := http.NewRequest("GET", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/content-compiler-cr/api/v1/content/amcn/amcplus/path")
      if strings.HasPrefix(u.path, "/movies/") {
         b.WriteString("/watch")
      }
      b.WriteString(u.path)
      return b.String()
   }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (a Auth) Playback(u URL) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Ad_Tags struct {
            Lat int `json:"lat"`
            Mode string `json:"mode"`
            PPID int `json:"ppid"`
            Player_Height int `json:"playerHeight"`
            Player_Width int `json:"playerWidth"`
            URL string `json:"url"`
         } `json:"adtags"`
      }
      s.Ad_Tags.Mode = "on-demand"
      s.Ad_Tags.URL = "-"
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
   req.URL.Path = "/playback-id/api/v1/playback/" + u.nid
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
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

func (a Auth) Refresh() (Raw_Auth, error) {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/refresh"
   req.Header.Set("Authorization", "Bearer " + a.Data.Refresh_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

type Raw_Auth []byte

func (a Auth) Login(email, password string) (Raw_Auth, error) {
   body, err := json.Marshal(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/login"
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
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
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

func (p Playback) HTTP_DASH() (*Source, error) {
   for _, s := range p.body.Data.Playback_JSON_Data.Sources {
      if strings.HasPrefix(s.Src, "https://") {
         if s.Type == "application/dash+xml" {
            return &s, nil
         }
      }
   }
   return nil, errors.New("data.playbackJsonData.sources")
}

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

type Playback struct {
   header http.Header
   body struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
}

func (p Playback) Request_Header() http.Header {
   return http.Header{
      "bcov-auth": {p.header.Get("X-AMCN-BC-JWT")},
   }
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (p Playback) Request_URL() (string, error) {
   v, err := p.HTTP_DASH()
   if err != nil {
      return "", err
   }
   return v.Key_Systems.Widevine.License_URL, nil
}

func (u URL) String() string {
   return u.path
}

// https://www.amcplus.com/movies/queen-of-earth--1026724
type URL struct {
   path string // /movies/queen-of-earth--1026724
   nid string // 1026724
}

func (u *URL) Set(s string) error {
   var found bool
   _, u.path, found = strings.Cut(s, "amcplus.com")
   if !found {
      return errors.New("amcplus.com")
   }
   _, u.nid, found = strings.Cut(s, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}
