package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

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
   v    *struct {
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
type OperationPlay struct {
   Data *struct {
      ArticleAssetPlay struct {
         Entitlements []struct {
            Manifest string
            Protocol string
         }
      }
   }
   raw []byte
}

/////////

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
      s.Variables.AssetId = asset.Id
      s.Variables.ArticleId = asset.article.Data.Article.Id
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
   var play OperationPlay
   play.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &play, nil
}

func (o *OperationPlay) Unmarshal() error {
   o.v = pointer(o.v)
   return json.Unmarshal(o.Data, o.v)
}

func (o OperationPlay) Dash() (string, bool) {
   for _, title := range o.v.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}
