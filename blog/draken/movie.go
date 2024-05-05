package draken

import (
   "bytes"
   "encoding/json"
   "net/http"
)

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

const magine_accesstoken = "22cc71a2-8b77-4819-95b0-8c90f4cf5663"

func new_movie(custom_id string) (*full_movie, error) {
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
      return nil, err
   }
   req.Header = http.Header{
      "magine-accesstoken": {magine_accesstoken},
      "x-forwarded-for": {"78.64.0.0"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Data struct {
         Viewer struct {
            ViewableCustomId struct {
               DefaultPlayable full_movie
            }
         }
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data.Viewer.ViewableCustomId.DefaultPlayable, nil
}

type full_movie struct {
   ID string
}
