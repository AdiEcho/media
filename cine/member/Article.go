package member

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

func (e Encoding) Year() int {
   if v, ok := e.Article.year(); ok {
      v, _ := strconv.Atoi(v)
      return v
   }
   return 0
}

func (d DataArticle) year() (string, bool) {
   for _, meta := range d.Metas {
      if meta.Key == "year" {
         return meta.Value, true
      }
   }
   return "", false
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

type Encoding struct {
   Article *DataArticle
   Play *AssetPlay
}

func (e Encoding) Marshal() ([]byte, error) {
   return json.MarshalIndent(e, "", " ")
}

func (d DataArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range d.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
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
   var s struct {
      Data struct {
         Article DataArticle
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   for _, asset := range s.Data.Article.Assets {
      asset.article = &s.Data.Article
   }
   return &s.Data.Article, nil
}

const query_article = `
query($articleUrlSlug: String) {
   Article(full_url_slug: $articleUrlSlug) {
      ... on Article {
         assets {
            ... on Asset {
               id
               linked_type
            }
         }
         canonical_title
         id
         metas(output: html) {
            ... on ArticleMeta {
               key
               value
            }
         }
      }
   }
}
`

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *DataArticle
}

// https://www.cinemember.nl/nl/films/american-hustle
func (a *ArticleSlug) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "cinemember.nl")
   s = strings.TrimPrefix(s, "/nl")
   s = strings.TrimPrefix(s, "/")
   *a = ArticleSlug(s)
   return nil
}

func (a ArticleSlug) String() string {
   return string(a)
}

type ArticleSlug string

func (Encoding) Show() string {
   return ""
}

func (Encoding) Season() int {
   return 0
}

func (Encoding) Episode() int {
   return 0
}

func (e Encoding) Title() string {
   return e.Article.CanonicalTitle
}

func (e *Encoding) Unmarshal(text []byte) error {
   return json.Unmarshal(text, e)
}
