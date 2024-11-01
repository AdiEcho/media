package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
   "strings"
)

func (*OperationArticle) Marshal(web *Address) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleUrlSlug string `json:"articleUrlSlug"`
      } `json:"variables"`
   }
   value.Variables.ArticleUrlSlug = web.Path
   value.Query = query_article
   body, err := json.Marshal(value)
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
   return io.ReadAll(resp.Body)
}

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *OperationArticle
}

// NO ANONYMOUS QUERY
const query_article = `
query Article($articleUrlSlug: String) {
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

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "cinemember.nl")
   s = strings.TrimPrefix(s, "/nl")
   a.Path = strings.TrimPrefix(s, "/")
   return nil
}

type Address struct {
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (o *OperationArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range o.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (o *OperationArticle) Title() string {
   return o.CanonicalTitle
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

func (*OperationArticle) Episode() int {
   return 0
}

func (*OperationArticle) Season() int {
   return 0
}

func (*OperationArticle) Show() string {
   return ""
}

type OperationArticle struct {
   Assets         []*ArticleAsset
   CanonicalTitle string `json:"canonical_title"`
   Id             int
   Metas          []struct {
      Key   string
      Value string
   }
}

func (o *OperationArticle) Unmarshal(data []byte) error {
   var value struct {
      Data struct {
         Article OperationArticle
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *o = value.Data.Article
   for _, asset := range o.Assets {
      asset.article = o
   }
   return nil
}
