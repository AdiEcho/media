package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
   "strings"
)

func (o *OperationArticle) Unmarshal() error {
   o.v = pointer(o.v)
   err := json.Unmarshal(o.Data, o.v)
   if err != nil {
      return err
   }
   for _, asset := range o.v.Assets {
      asset.article = o
   }
   return nil
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

const query_user = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

func pointer[T any](value *T) *T {
   return new(T)
}

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

func (OperationArticle) Episode() int {
   return 0
}

func (OperationArticle) Season() int {
   return 0
}

func (OperationArticle) Show() string {
   return ""
}

func (o OperationArticle) Title() string {
   return o.v.CanonicalTitle
}

func (o OperationArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range o.v.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (o OperationArticle) Year() int {
   for _, meta := range o.v.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (o *OperationUser) New(email, password string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            Email    string `json:"email"`
            Password string `json:"password"`
         } `json:"variables"`
      }
      s.Query = query_user
      s.Variables.Email = email
      s.Variables.Password = password
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   o.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *OperationArticle
}

type OperationUser struct {
   Data []byte
   v *struct {
      Data struct {
         UserAuthenticate struct {
            AccessToken string `json:"access_token"`
         }
      }
   }
}

func (o *OperationUser) Unmarshal() error {
   o.v = pointer(o.v)
   return json.Unmarshal(o.Data, o.v)
}

type OperationArticle struct {
   Data []byte
   v *struct {
      Assets         []*ArticleAsset
      CanonicalTitle string `json:"canonical_title"`
      Id             int
      Metas          []struct {
         Key   string
         Value string
      }
   }
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
   var data struct {
      Data struct {
         Article json.RawMessage
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &OperationArticle{Data: data.Data.Article}, nil
}
