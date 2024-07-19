package draken

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

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
   var data struct {
      Data struct {
         Viewer struct {
            ViewableCustomId *FullMovie
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   if v := data.Data.Viewer.ViewableCustomId; v != nil {
      return v, nil
   }
   return nil, FullMovie{}
}

func (FullMovie) Error() string {
   return "FullMovie"
}

type FullMovie struct {
   DefaultPlayable struct {
      Id string
   }
   ProductionYear int `json:",string"`
   Title          string
}

type Namer struct {
   Movie *FullMovie
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
   return n.Movie.Title
}

func (n Namer) Year() int {
   return n.Movie.ProductionYear
}

const get_custom_id = `
query($customId: ID!) {
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
