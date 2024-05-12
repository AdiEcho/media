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
         id
         assets {
            ... on Asset {
               id
               linked_type
            }
         }
      }
   }
}
`

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

type data_article struct {
   ID int
   Assets []*article_asset
}

type article_asset struct {
   ID int
   Linked_Type string
   article *data_article
}
