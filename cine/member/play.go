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

func (o OperationPlay) Dash() (string, bool) {
   for _, title := range o.v.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
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
