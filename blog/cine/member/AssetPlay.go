package member

import (
   "bytes"
   "encoding/json"
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

func (a authenticate) play(asset article_asset) (*asset_play, error) {
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
      "authorization": {"Bearer " + a.Data.UserAuthenticate.AccessToken},
      "content-type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(asset_play)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type asset_play struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []struct {
            Manifest string
            Protocol string
         }
      }
   }
}
