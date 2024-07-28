package criterion

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a AuthToken) Files(item *EmbedItem) (VideoFiles, error) {
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

func (a AuthToken) Video(slug string) (*EmbedItem, error) {
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
   item := new(EmbedItem)
   err = json.NewDecoder(resp.Body).Decode(item)
   if err != nil {
      return nil, err
   }
   return item, nil
}

func (a *AuthToken) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

type AuthToken struct {
   AccessToken string `json:"access_token"`
}

func NewAuthToken(username, password string) ([]byte, error) {
   resp, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id":  {client_id},
      "grant_type": {"password"},
      "password":   {password},
      "username":   {username},
   })
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
