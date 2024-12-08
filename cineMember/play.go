package cineMember

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

type OperationPlay struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []Entitlement
      }
   }
   Errors []struct {
      Message string
   }
}

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

func (o *OperationPlay) Dash() (*Entitlement, bool) {
   for _, title := range o.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return &title, true
      }
   }
   return nil, false
}

func (e *Entitlement) RequestUrl() (string, bool) {
   return e.KeyDeliveryUrl, true
}

func (*Entitlement) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (*Entitlement) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (*Entitlement) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type Entitlement struct {
   KeyDeliveryUrl string `json:"key_delivery_url"`
   Manifest string
   Protocol string
}

const query_play = `
mutation($article_id: Int, $asset_id: Int) {
   ArticleAssetPlay(article_id: $article_id asset_id: $asset_id) {
      entitlements {
         ... on ArticleAssetPlayEntitlement {
            key_delivery_url
            manifest
            protocol
         }
      }
   }
}
`
