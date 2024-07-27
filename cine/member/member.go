package member

import (
   "bytes"
   "encoding/json"
   "errors"
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

const query_play = `
mutation($article_id: Int, $asset_id: Int) {
   ArticleAssetPlay(article_id: $article_id asset_id: $asset_id) {
      entitlements {
         ... on ArticleAssetPlayEntitlement {
            manifest
            protocol
         }
      }
   }
}
`

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

func (o OperationPlay) Dash() (string, bool) {
   for _, title := range o.v.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}

// geo block, not x-forwarded-for
func (o OperationUser) Play(asset *ArticleAsset) (*OperationPlay, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            ArticleId int `json:"article_id"`
            AssetId   int `json:"asset_id"`
         } `json:"variables"`
      }
      s.Query = query_play
      s.Variables.ArticleId = asset.article.v.Id
      s.Variables.AssetId = asset.Id
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.audienceplayer.com/graphql/2/user",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + o.v.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var data struct {
      Data struct {
         ArticleAssetPlay json.RawMessage
      }
   }
   err = json.Unmarshal(text, &data)
   if err != nil {
      return nil, err
   }
   if v := data.Data.ArticleAssetPlay; v != nil {
      return &OperationPlay{Data: v}, nil
   }
   return nil, errors.New(string(text))
}

func (o *OperationUser) Unmarshal() error {
   o.v = pointer(o.v)
   return json.Unmarshal(o.Data, o.v)
}

func (o *OperationPlay) Unmarshal() error {
   o.v = pointer(o.v)
   return json.Unmarshal(o.Data, o.v)
}

type OperationPlay struct {
   Data []byte
   v *struct {
      Entitlements []struct {
         Manifest string
         Protocol string
      }
   }
}
