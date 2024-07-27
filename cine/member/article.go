package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

func (o *OperationArticle) SetRaw(raw []byte) {
   o.raw = raw
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

func (OperationArticle) Episode() int {
   return 0
}

func (OperationArticle) Season() int {
   return 0
}

func (OperationArticle) Show() string {
   return ""
}

func (o *OperationArticle) Unmarshal() error {
   o.Data = pointer(o.Data)
   err := json.Unmarshal(o.raw, o)
   if err != nil {
      return err
   }
   for _, asset := range o.Data.Article.Assets {
      asset.article = o
   }
   return nil
}

func (o OperationArticle) Title() string {
   return o.Data.Article.CanonicalTitle
}

func (o OperationArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range o.Data.Article.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (o OperationArticle) Year() int {
   for _, meta := range o.Data.Article.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (a ArticleSlug) Article() (*OperationArticle, error) {
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
   var article OperationArticle
   article.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &article, nil
}

type OperationArticle struct {
   Data *struct {
      Article struct {
         Assets         []*ArticleAsset
         CanonicalTitle string `json:"canonical_title"`
         Id             int
         Metas          []struct {
            Key   string
            Value string
         }
      }
   }
   raw []byte
}
