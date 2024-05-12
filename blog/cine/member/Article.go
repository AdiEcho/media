package member

import (
   "bytes"
   "encoding/json"
   "net/http"
)

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

func (d data_article) film() (*article_asset, bool) {
   for _, asset := range d.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

type article_asset struct {
   ID int
   LinkedType string `json:"linked_type"`
   article *data_article
}

func new_article(slug string) (*data_article, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            ArticleUrlSlug string `json:"articleUrlSlug"`
         } `json:"variables"`
      }
      s.Query = query_article
      s.Variables.ArticleUrlSlug = slug
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Data struct {
         Article data_article
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   for _, asset := range s.Data.Article.Assets {
      asset.article = &s.Data.Article
   }
   return &s.Data.Article, nil
}

func (data_article) Show() string {
   return ""
}

func (data_article) Season() int {
   return 0
}

func (data_article) Episode() int {
   return 0
}

func (d data_article) Title() string {
   return d.CanonicalTitle
}

type data_article struct {
   Assets []*article_asset
   CanonicalTitle string `json:"canonical_title"`
   ID int
   Metas []struct {
      Key string
      Value string
   }
}

func (data_article) Year() int {
   return 0
}
