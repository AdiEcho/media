package member

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

func (o *OperationPlay) Dash() (string, bool) {
   for _, title := range o.Data.ArticleAssetPlay.Entitlements {
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

// hard geo block
func (o *OperationUser) Play(
   asset *ArticleAsset, data *[]byte,
) (*OperationPlay, error) {
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
   body, err := json.Marshal(value)
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
      "authorization": {"Bearer " + o.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   body, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if data != nil {
      *data = body
      return nil, nil
   }
   var play OperationPlay
   err = play.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   return &play, nil
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
