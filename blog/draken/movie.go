package draken

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type full_movie struct {
   Data struct {
      Viewer struct {
         ViewableCustomId struct {
            DefaultPlayable struct {
               ID string
            }
         }
      }
   }
}

const get_custom_id = `
query GetCustomIdFullMovie($customId: ID!) {
   viewer {
      viewableCustomId(customId: $customId) {
         ... on Movie {
            defaultPlayable {
               id
            }
         }
      }
   }
}
`

func (f *full_movie) New(custom_id string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            CustomId string `json:"customId"`
         } `json:"variables"`
      }
      s.Query = get_custom_id
      s.Variables.CustomId = custom_id
      return json.Marshal(s)
   }()
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/apiql/v2",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "magine-accesstoken": {"22cc71a2-8b77-4819-95b0-8c90f4cf5663"},
      "x-forwarded-for": {"78.64.0.0"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(f)
}
