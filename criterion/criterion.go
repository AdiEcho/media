package criterion

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

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
   req.Header.Set("authorization", "Bearer "+a.V.AccessToken)
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
   item := new(EmbedItem)
   err = json.NewDecoder(res.Body).Decode(item)
   if err != nil {
      return nil, err
   }
   return item, nil
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

func (EmbedItem) Episode() int {
   return 0
}

func (EmbedItem) Season() int {
   return 0
}

func (EmbedItem) Show() string {
   return ""
}

func (e EmbedItem) Title() string {
   return e.Name
}

func (e EmbedItem) Year() int {
   return e.Metadata.YearReleased
}
const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

func (a *AuthToken) New(username, password string) error {
   res, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id":  {client_id},
      "grant_type": {"password"},
      "password":   {password},
      "username":   {username},
   })
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

type AuthToken struct {
   Data []byte
   V    struct {
      AccessToken string `json:"access_token"`
   }
}

func (a *AuthToken) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}
