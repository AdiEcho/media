package criterion

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a AuthToken) item(slug string) (*embed_item, error) {
   address := func() string {
      var b strings.Builder
      b.WriteString("https://api.vhx.com/collections/")
      b.WriteString(slug)
      b.WriteString("/items?site_id=59054")
      return b.String()
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer "+a.v.AccessToken)
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
   var s struct {
      Embedded struct {
         Items []embed_item
      } `json:"_embedded"`
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Embedded.Items[0], nil
}

const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

type AuthToken struct {
   data []byte
   v    struct {
      AccessToken string `json:"access_token"`
   }
}

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
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AuthToken) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}

type embed_item struct {
   ID       int64
   Metadata struct {
      YearReleased int `json:"year_released"`
   }
   Name string
}

func (embed_item) Episode() int {
   return 0
}

func (embed_item) Season() int {
   return 0
}

func (embed_item) Show() string {
   return ""
}

func (e embed_item) Title() string {
   return e.Name
}

func (e embed_item) Year() int {
   return e.Metadata.YearReleased
}
