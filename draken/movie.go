package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

const get_custom_id = `
query GetCustomIdFullMovie($customId: ID!) {
   viewer {
      viewableCustomId(customId: $customId) {
         ... on Movie {
            defaultPlayable {
               id
            }
            productionYear
            title
         }
      }
   }
}
`

func graphql_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

type FullMovie struct {
   DefaultPlayable struct {
      ID string
   }
   ProductionYear int `json:",string"`
   Title          string
}

func NewMovie(custom_id string) (*FullMovie, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            CustomId string `json:"customId"`
         } `json:"variables"`
      }
      s.Variables.CustomId = custom_id
      s.Query = graphql_compact(get_custom_id)
      return json.MarshalIndent(s, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/apiql/v2",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   magine_accesstoken.set(req.Header)
   x_forwarded_for.set(req.Header)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var s struct {
      Data struct {
         Viewer struct {
            ViewableCustomId *FullMovie
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   if v := s.Data.Viewer.ViewableCustomId; v != nil {
      return v, nil
   }
   return nil, errors.New(`"viewableCustomId": null`)
}

type Namer struct {
   F *FullMovie
}

func (Namer) Episode() int {
   return 0
}

func (Namer) Season() int {
   return 0
}

func (Namer) Show() string {
   return ""
}

func (n Namer) Title() string {
   return n.F.Title
}

func (n Namer) Year() int {
   return n.F.ProductionYear
}
