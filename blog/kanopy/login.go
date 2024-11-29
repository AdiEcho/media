package kanopy

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

// good for 10 years
type web_token struct {
   Jwt string
   UserId int64
}

func (w *web_token) videos(id int64) (*videos_response, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/videos/" + strconv.FormatInt(id, 10)
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   videos := &videos_response{}
   err = json.NewDecoder(resp.Body).Decode(videos)
   if err != nil {
      return nil, err
   }
   return videos, nil
}

func (*videos_response) Show() string {
   return ""
}

func (*videos_response) Season() int {
   return 0
}

func (*videos_response) Episode() int {
   return 0
}

func (v *videos_response) Title() string {
   return v.Video.Title
}

type videos_response struct {
   Video struct {
      ProductionYear int
      Title string
   }
}

func (v *videos_response) Year() int {
   return v.Video.ProductionYear
}

func (web_token) marshal(email, password string) ([]byte, error) {
   value := map[string]any{
      "credentialType": "email",
      "emailUser": map[string]string{
         "email": email,
         "password": password,
      },
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/login", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {user_agent},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (w *web_token) unmarshal(data []byte) error {
   return json.Unmarshal(data, w)
}

func (*poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (*poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (p *poster) RequestHeader() (http.Header, error) {
   h := http.Header{
      "authorization": {"Bearer " + p.token.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   return h, nil
}

type poster struct {
   manifest *video_manifest
   token *web_token
}

func (p *poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Scheme = "https"
   u.Host = "www.kanopy.com"
   u.Path = "/kapi/licenses/widevine/" + p.manifest.DrmLicenseId
   return u.String(), true
}
