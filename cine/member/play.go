package member

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

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

// hard geo block
func (OperationPlay) Marshal(
   user *OperationUser, asset *ArticleAsset,
) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleId int `json:"article_id"`
         AssetId   int `json:"asset_id"`
      } `json:"variables"`
   }
   value.Query = query_play
   value.Variables.AssetId = asset.Id
   value.Variables.ArticleId = asset.article.Id
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.audienceplayer.com/graphql/2/user",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + user.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (o *OperationPlay) Dash() (string, bool) {
   for _, title := range o.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}

type OperationPlay struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []struct {
            Manifest string
            Protocol string
         }
      }
   }
   Errors []struct {
      Message string
   }
}

func (o *OperationPlay) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, o)
   if err != nil {
      return err
   }
   if v := o.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   return nil
}
