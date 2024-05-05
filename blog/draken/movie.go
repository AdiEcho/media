package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

func new_movie(custom_id string) (*full_movie, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            CustomId string `json:"customId"`
         } `json:"variables"`
      }
      s.Variables.CustomId = custom_id
      s.Query = graphql_compact(get_custom_id)
      return json.MarshalIndent(s, "", " ")
   }()
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/apiql/v2",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   magine_accesstoken.set(req.Header)
   x_forwarded_for.set(req.Header)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Data struct {
         Viewer struct {
            ViewableCustomId *struct {
               DefaultPlayable full_movie
            }
         }
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   if v := s.Data.Viewer.ViewableCustomId; v != nil {
      return &v.DefaultPlayable, nil
   }
   return nil, errors.New(`"viewableCustomId": null`)
}
func graphql_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
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

type full_movie struct {
   ID string
}
