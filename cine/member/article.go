package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
   "strings"
)

func (ArticleAsset) Error() string {
   return "ArticleAsset"
}

type ArticleSlug string

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

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *OperationArticle
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

func (o *OperationArticle) Title() string {
   return o.CanonicalTitle
}

func (o *OperationArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range o.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (o *OperationArticle) Year() int {
   for _, meta := range o.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (a ArticleSlug) Article() (*OperationArticle, error) {
   var body struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleUrlSlug ArticleSlug `json:"articleUrlSlug"`
      } `json:"variables"`
   }
   body.Variables.ArticleUrlSlug = a
   body.Query = query_article
   raw, err := json.Marshal(body)
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(raw),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var article OperationArticle
   article.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &article, nil
}

type OperationArticle struct {
   Assets         []*ArticleAsset
   CanonicalTitle string `json:"canonical_title"`
   Id             int
   Metas          []struct {
      Key   string
      Value string
   }
   Raw []byte `json:"-"`
}

func (o *OperationArticle) Unmarshal() error {
   var body struct {
      Article OperationArticle
   }
   err := json.Unmarshal(o.Raw, &body)
   if err != nil {
      return err
   }
   *o = body.Article
   for _, asset := range o.Assets {
      asset.article = o
   }
   return nil
}
