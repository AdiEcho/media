package member

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

func (o *OperationPlay) Dash() (string, bool) {
   for _, title := range o.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
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
   Entitlements []struct {
      Manifest string
      Protocol string
   }
   Raw []byte `json:"-"`
}

func (o *OperationPlay) Unmarshal() error {
   var value struct {
      Data struct {
         ArticleAssetPlay OperationPlay
      }
      Errors []struct {
         Message string
      }
   }
   err := json.Unmarshal(o.Raw, &value)
   if err != nil {
      return err
   }
   if v := value.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   *o = value.Data.ArticleAssetPlay
   return nil
}

// geo block, not x-forwarded-for
func (o *OperationUser) Play(asset *ArticleAsset) (*OperationPlay, error) {
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
      "authorization": {"Bearer " + o.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var play OperationPlay
   play.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &play, nil
}
