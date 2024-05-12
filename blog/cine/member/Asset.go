package member

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

// geo block - VPN not x-forwarded-for
func (a authenticate) play(asset *article_asset) (*asset_play, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            ArticleId int `json:"article_id"`
            AssetId int `json:"asset_id"`
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
      "authorization": {"Bearer " + a.v.Data.UserAuthenticate.AccessToken},
      "content-type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var s struct {
      Data struct {
         ArticleAssetPlay *asset_play
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

type asset_play struct {
   Entitlements []struct {
      Manifest string
      Protocol string
   }
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
