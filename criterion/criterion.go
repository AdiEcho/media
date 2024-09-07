package criterion

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (*VideoFiles) DashError() error {
   return errors.New("VideoFiles.Dash")
}

func (v VideoFiles) Dash() (*VideoFile, bool) {
   for _, file := range v {
      if file.Method == "dash" {
         return &file, true
      }
   }
   return nil, false
}

func (a *AuthToken) New(username, password string) error {
   resp, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id":  {client_id},
      "grant_type": {"password"},
      "password":   {password},
      "username":   {username},
   })
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AuthToken) Files(item *EmbedItem) (VideoFiles, error) {
   req, err := http.NewRequest("", item.Links.Files.Href, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer "+a.AccessToken)
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
   var files VideoFiles
   err = json.NewDecoder(resp.Body).Decode(&files)
   if err != nil {
      return nil, err
   }
   return files, nil
}

func (a *AuthToken) Video(slug string) (*EmbedItem, error) {
   req, err := http.NewRequest("", "https://api.vhx.com/videos", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteByte('/')
      b.WriteString(slug)
      b.WriteString("?url=")
      b.WriteString(slug)
      return b.String()
   }()
   req.Header.Set("authorization", "Bearer "+a.AccessToken)
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
   item := &EmbedItem{}
   err = json.NewDecoder(resp.Body).Decode(item)
   if err != nil {
      return nil, err
   }
   return item, nil
}

type AuthToken struct {
   AccessToken string `json:"access_token"`
   Raw []byte `json:"-"`
}

func (a *AuthToken) Unmarshal() error {
   return json.Unmarshal(a.Raw, a)
}

type VideoFile struct {
   DrmAuthorizationToken string `json:"drm_authorization_token"`
   Links                 struct {
      Source struct {
         Href string
      }
   } `json:"_links"`
   Method string
}

type VideoFiles []VideoFile

func (*VideoFile) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (*VideoFile) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (*VideoFile) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

type EmbedItem struct {
   Links struct {
      Files struct {
         Href string
      }
   } `json:"_links"`
   Metadata struct {
      YearReleased int `json:"year_released"`
   }
   Name string
}

const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

func (*EmbedItem) Episode() int {
   return 0
}

func (*EmbedItem) Season() int {
   return 0
}

func (*EmbedItem) Show() string {
   return ""
}

func (e *EmbedItem) Title() string {
   return e.Name
}

func (e *EmbedItem) Year() int {
   return e.Metadata.YearReleased
}

func (v *VideoFile) RequestUrl() (string, bool) {
   b := []byte("https://drm.vhx.com/v2/widevine?token=")
   b = append(b, v.DrmAuthorizationToken...)
   return string(b), true
}
