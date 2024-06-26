package member

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

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
      s.Variables.ArticleId = asset.article.ID
      s.Variables.AssetId = asset.ID
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
      "authorization": {"Bearer " + a.V.Data.UserAuthenticate.AccessToken},
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
   var s struct {
      Data struct {
         ArticleAssetPlay *AssetPlay
      }
   }
   err = json.Unmarshal(text, &s)
   if err != nil {
      return nil, err
   }
   if v := s.Data.ArticleAssetPlay; v != nil {
      return v, nil
   }
   return nil, errors.New(string(text))
}

type AssetPlay struct {
   Entitlements []struct {
      Manifest string
      Protocol string
   }
}

func (a AssetPlay) DASH() (string, bool) {
   for _, title := range a.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}

type Authenticate struct {
   Data []byte
   V    struct {
      Data struct {
         UserAuthenticate struct {
            AccessToken string `json:"access_token"`
         }
      }
   }
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

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

const user_authenticate = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`
