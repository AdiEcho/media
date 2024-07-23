package member

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

func (e DataArticle) Title() string {
   return e.Article.CanonicalTitle
}

func (e DataArticle) Year() int {
   if v, ok := e.Article.year(); ok {
      v, _ := strconv.Atoi(v)
      return v
   }
   return 0
}

func (DataArticle) Episode() int {
   return 0
}

func (DataArticle) Season() int {
   return 0
}

func (d DataArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range d.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (d DataArticle) year() (string, bool) {
   for _, meta := range d.Metas {
      if meta.Key == "year" {
         return meta.Value, true
      }
   }
   return "", false
}

func (DataArticle) Show() string {
   return ""
}

type DataArticle struct {
   Assets         []*ArticleAsset
   CanonicalTitle string `json:"canonical_title"`
   Id             int
   Metas          []struct {
      Key   string
      Value string
   }
}

func (e DataArticle) Marshal() ([]byte, error) {
   return json.MarshalIndent(e, "", " ")
}

func (e *DataArticle) Unmarshal(text []byte) error {
   return json.Unmarshal(text, e)
}

func (a ArticleSlug) Article() (*DataArticle, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            ArticleUrlSlug ArticleSlug `json:"articleUrlSlug"`
         } `json:"variables"`
      }
      s.Variables.ArticleUrlSlug = a
      s.Query = query_article
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var data struct {
      Data struct {
         Article DataArticle
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   for _, asset := range data.Data.Article.Assets {
      asset.article = &data.Data.Article
   }
   return &data.Data.Article, nil
}
