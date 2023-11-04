package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
   "time"
)

func New_Metadata(guid int64) (*Metadata, error) {
   body, err := func() ([]byte, error) {
      var p page_request
      p.Variables.Name = strconv.FormatInt(guid, 10)
      p.Query = graphQL_compact(query)
      p.Variables.App = "nbc"
      p.Variables.One_App = true
      p.Variables.Platform = "android"
      p.Variables.Type = "VIDEO"
      return json.MarshalIndent(p, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var page struct {
      Data struct {
         Bonanza_Page struct {
            Metadata *Metadata
         } `json:"bonanzaPage"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   if page.Data.Bonanza_Page.Metadata == nil {
      return nil, errors.New(".data.bonanzaPage.metadata")
   }
   return page.Data.Bonanza_Page.Metadata, nil
}

type Metadata struct {
   Air_Date string `json:"airDate"`
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Secondary_Title string `json:"secondaryTitle"`
   Series_Short_Title *string `json:"seriesShortTitle"`
   Season_Number *int64 `json:"seasonNumber,string"`
   Episode_Number *int64 `json:"episodeNumber,string"`
}

func (m Metadata) Series() string {
   return *m.Series_Short_Title
}

func (m Metadata) Season() (int64, error) {
   return *m.Season_Number, nil
}

func (m Metadata) Episode() (int64, error) {
   return *m.Episode_Number, nil
}

func (m Metadata) Title() string {
   return m.Secondary_Title
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, m.Air_Date)
}

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      oneApp: $oneApp
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            airDate
            episodeNumber
            mpxAccountId
            mpxGuid
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphQL_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      One_App bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      User_ID string `json:"userId"`
   } `json:"variables"`
}
