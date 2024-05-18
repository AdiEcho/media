package criterion

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (a auth_token) video() (*site_video, error) {
   req, err := http.NewRequest("", "https://api.vhx.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v2/sites/59054/videos/455774"
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   video := new(site_video)
   err = json.NewDecoder(res.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (namer) Show() string {
   return ""
}

func (namer) Season() int {
   return 0
}

func (namer) Episode() int {
   return 0
}

type namer struct {
   s *site_video
}

func (n namer) Title() string {
   return n.s.Title
}

type site_video struct {
   Metadata struct {
      Custom []struct {
         Key string
         Value string
      }
   }
   Title string
}

func (s site_video) year_released() (string, bool) {
   for _, custom := range s.Metadata.Custom {
      if custom.Key == "year_released" {
         return custom.Value, true
      }
   }
   return "", false
}

func (n namer) Year() int {
   if v, ok := n.s.year_released(); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}
const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

type auth_token struct {
   data []byte
   v struct {
      AccessToken string `json:"access_token"`
   }
}

func (a *auth_token) New(username, password string) error {
   res, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id": {client_id},
      "grant_type": {"password"},
      "password": {password},
      "username": {username},
   })
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *auth_token) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
