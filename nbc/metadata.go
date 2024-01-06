package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
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
   var s struct {
      Data struct {
         Bonanza_Page struct {
            Metadata Metadata
         } `json:"bonanzaPage"`
      }
      Errors []struct {
         Message string
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   if len(s.Errors) >= 1 {
      return nil, errors.New(s.Errors[0].Message)
   }
   return &s.Data.Bonanza_Page.Metadata, nil
}

type Metadata struct {
   Air_Date string `json:"airDate"`
   Episode_Number string `json:"episodeNumber"`
   MPX_Account_ID int64 `json:"mpxAccountId,string"`
   MPX_GUID int64 `json:"mpxGuid,string"`
   Programming_Type string `json:"programmingType"`
   Season_Number string `json:"seasonNumber"`
   Secondary_Title string `json:"secondaryTitle"`
   Series_Short_Title string `json:"seriesShortTitle"`
}

func (Metadata) Owner() (string, bool) {
   return "", false
}

func (m Metadata) Season() (string, bool) {
   return m.Season_Number, m.Season_Number != ""
}

func (m Metadata) Episode() (string, bool) {
   return m.Episode_Number, m.Episode_Number != ""
}

func (m Metadata) Title() (string, bool) {
   return m.Secondary_Title, true
}

func (m Metadata) Show() (string, bool) {
   return m.Series_Short_Title, m.Series_Short_Title != ""
}

func (m Metadata) Year() (string, bool) {
   if m.Series_Short_Title != "" {
      return "", false
   }
   year, _, _ := strings.Cut(m.Air_Date, "-")
   return year, true
}
