package amc

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

func Unauth() (*Auth_ID, error) {
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
   var auth Auth_ID
   auth.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a Auth_ID) Content(u URL) (*Content, error) {
   req, err := http.NewRequest("GET", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Value.Data.Access_Token},
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
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (a Auth_ID) Playback(u URL) (*Playback, error) {
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
      "Authorization": {"Bearer " + a.Value.Data.Access_Token},
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
   var play Playback
   play.header = res.Header
   if err := json.NewDecoder(res.Body).Decode(&play.body); err != nil {
      return nil, err
   }
   return &play, nil
}

func (a *Auth_ID) Unmarshal() error {
   return json.Unmarshal(a.Raw, a.Value)
}

func (a *Auth_ID) Refresh() error {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/refresh"
   req.Header.Set("Authorization", "Bearer " + a.Value.Data.Refresh_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.Value = nil
   return nil
}

func (a *Auth_ID) Login(email, password string) error {
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
      "Authorization": {"Bearer " + a.Value.Data.Access_Token},
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
   a.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.Value = nil
   return nil
}

type Auth_ID struct {
   Raw []byte
   Value *struct {
      Data struct {
         Access_Token string
         Refresh_Token string
      }
   }
}
