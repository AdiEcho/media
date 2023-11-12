package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

type On_Demand struct {
   Playback_URL string `json:"playbackUrl"`
}

func (m Metadata) On_Demand() (*On_Demand, error) {
   req, err := http.NewRequest("GET", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MPX_Account_ID, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MPX_GUID, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {m.Programming_Type},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   video := new(On_Demand)
   if err := json.NewDecoder(res.Body).Decode(video); err != nil {
      return nil, err
   }
   return video, nil
}

type Metadata struct {
   Air_Date string `json:"airDate"`
   Episode_Number int64 `json:"episodeNumber,string"`
   MPX_Account_ID int64 `json:"mpxAccountId,string"`
   MPX_GUID int64 `json:"mpxGuid,string"`
   Programming_Type string `json:"programmingType"`
   Season_Number int64 `json:"seasonNumber,string"`
   Secondary_Title string `json:"secondaryTitle"`
   Series_Short_Title string `json:"seriesShortTitle"`
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
            programmingType
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

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
   var s struct {
      Data struct {
         Bonanza_Page struct {
            Metadata Metadata
         } `json:"bonanzaPage"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.Data.Bonanza_Page.Metadata, nil
}

func (m Metadata) Series() string {
   return m.Series_Short_Title
}

func (m Metadata) Season() (int64, error) {
   return m.Season_Number, nil
}

func (m Metadata) Episode() (int64, error) {
   return m.Episode_Number, nil
}

func (m Metadata) Title() string {
   return m.Secondary_Title
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, m.Air_Date)
}

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
