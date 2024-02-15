package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

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

func (m Metadata) Title() (string, bool) {
   return m.SecondaryTitle, true
}

func (m Metadata) Season() (string, bool) {
   if v := m.SeasonNumber; v != "" {
      return v, true
   }
   return "", false
}

func (m Metadata) Episode() (string, bool) {
   if v := m.EpisodeNumber; v != "" {
      return v, true
   }
   return "", false
}

func (m Metadata) Show() (string, bool) {
   if v := m.SeriesShortTitle; v != "" {
      return v, true
   }
   return "", false
}

func (m Metadata) Year() (string, bool) {
   if m.SeriesShortTitle != "" {
      return "", false
   }
   year, _, _ := strings.Cut(m.AirDate, "-")
   return year, true
}

func (m *Metadata) New(guid int) error {
   body, err := func() ([]byte, error) {
      var p page_request
      p.Variables.Name = strconv.Itoa(guid)
      p.Query = graphql_compact(query)
      p.Variables.App = "nbc"
      p.Variables.OneApp = true
      p.Variables.Platform = "android"
      p.Variables.Type = "VIDEO"
      return json.Marshal(p)
   }()
   if err != nil {
      return err
   }
   res, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
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
      return err
   }
   if len(s.Errors) >= 1 {
      return errors.New(s.Errors[0].Message)
   }
   *m = s.Data.BonanzaPage.Metadata
   return nil
}
