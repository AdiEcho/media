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

func (d *DataArticle) Unmarshal() error {
   d.v = pointer(d.v)
   err := json.Unmarshal(d.Data, d.v)
   if err != nil {
      return err
   }
   for _, asset := range d.v.Assets {
      asset.article = d
   }
   return nil
}

func (DataArticle) Episode() int {
   return 0
}

func (DataArticle) Season() int {
   return 0
}

func (DataArticle) Show() string {
   return ""
}

func (d DataArticle) Title() string {
   return d.v.CanonicalTitle
}

func (d DataArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range d.v.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (d DataArticle) Year() int {
   for _, meta := range d.v.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (ArticleAsset) Error() string {
   return "ArticleAsset"
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

const query_asset = `
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

func (a AssetPlay) Dash() (string, bool) {
   for _, title := range a.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}

// geo block - VPN not x-forwarded-for
func (a Authenticate) Play(asset *ArticleAsset) (*AssetPlay, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            ArticleId int `json:"article_id"`
            AssetId   int `json:"asset_id"`
         } `json:"variables"`
      }
      s.Query = query_asset
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
      "authorization": {"Bearer " + a.v.Data.UserAuthenticate.AccessToken},
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
         ArticleAssetPlay *AssetPlay
      }
   }
   err = json.Unmarshal(text, &data)
   if err != nil {
      return nil, err
   }
   if v := data.Data.ArticleAssetPlay; v != nil {
      return v, nil
   }
   return nil, errors.New(string(text))
}

func pointer[T any](value *T) *T {
   return new(T)
}

func (a *Authenticate) Unmarshal() error {
   a.v = pointer(a.v)
   return json.Unmarshal(a.Data, a.v)
}

func (a *Authenticate) New(email, password string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            Email    string `json:"email"`
            Password string `json:"password"`
         } `json:"variables"`
      }
      s.Query = user_authenticate
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
   a.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

const user_authenticate = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

type Authenticate struct {
   Data []byte
   v *struct {
      Data struct {
         UserAuthenticate struct {
            AccessToken string `json:"access_token"`
         }
      }
   }
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
         Article json.RawMessage
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &DataArticle{Data: data.Data.Article}, nil
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

/////////////

type DataArticle struct {
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

type AssetPlay struct {
   Entitlements []struct {
      Manifest string
      Protocol string
   }
}

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *DataArticle
}

type ArticleSlug string
