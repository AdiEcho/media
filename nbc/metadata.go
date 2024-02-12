package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

func NewMetadata(guid int64) (*Metadata, error) {
   body, err := func() ([]byte, error) {
      var p page_request
      p.Variables.Name = strconv.FormatInt(guid, 10)
      p.Query = graphql_compact(query)
      p.Variables.App = "nbc"
      p.Variables.OneApp = true
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
         BonanzaPage struct {
            Metadata Metadata
         }
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
   return &s.Data.BonanzaPage.Metadata, nil
}

type Metadata struct {
   AirDate string
   EpisodeNumber string
   MpxAccountId int64 `json:",string"`
   MpxGuid int64 `json:",string"`
   ProgrammingType string
   SeasonNumber string
   SecondaryTitle string
   SeriesShortTitle string
}

func (Metadata) Owner() (string, bool) {
   return "", false
}

func (m Metadata) Season() (string, bool) {
   return m.SeasonNumber, m.SeasonNumber != ""
}

func (m Metadata) Episode() (string, bool) {
   return m.EpisodeNumber, m.EpisodeNumber != ""
}

func (m Metadata) Title() (string, bool) {
   return m.SecondaryTitle, true
}

func (m Metadata) Show() (string, bool) {
   return m.SeriesShortTitle, m.SeriesShortTitle != ""
}

func (m Metadata) Year() (string, bool) {
   if m.SeriesShortTitle != "" {
      return "", false
   }
   year, _, _ := strings.Cut(m.AirDate, "-")
   return year, true
}
