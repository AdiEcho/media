package criterion

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type embed_item struct {
   Name string
   Metadata struct {
      YearReleased int `json:"year_released"`
   }
   ID       int64
}

func (a AuthToken) video(slug string) (*embed_item, error) {
   req, err := http.NewRequest("", "https://api.vhx.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/videos/" + slug
   req.URL.RawQuery = "url=" + slug
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
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
   item := new(embed_item)
   err = json.NewDecoder(res.Body).Decode(item)
   if err != nil {
      return nil, err
   }
   return item, nil
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
