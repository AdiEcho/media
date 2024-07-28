package member

import (
   "bytes"
   "encoding/json"
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

func (o *OperationPlay) Unmarshal() error {
   return json.Unmarshal(o.raw, o)
}

func (o OperationPlay) Dash() (string, bool) {
   for _, title := range o.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}

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
      "authorization": {"Bearer " + o.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var play OperationPlay
   play.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &play, nil
}

func (o *OperationPlay) SetRaw(raw []byte) {
   o.raw = raw
}

func (o OperationPlay) GetRaw() []byte {
   return o.raw
}
