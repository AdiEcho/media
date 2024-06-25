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

func (m Metadata) Title() string {
   if v := m.MovieShortTitle; v != "" {
      return v
   }
   return m.SecondaryTitle
}

type Metadata struct {
   AirDate time.Time
   EpisodeNumber int `json:",string"`
   MovieShortTitle string
   MpxAccountId int64 `json:",string"`
   MpxGuid int64 `json:",string"`
   ProgrammingType string
   SeasonNumber int `json:",string"`
   SecondaryTitle string
   SeriesShortTitle string
}

func (m Metadata) Year() int {
   return m.AirDate.Year()
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
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
   resp, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
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
   err = json.NewDecoder(resp.Body).Decode(&s)
   if err != nil {
      return err
   }
   if len(s.Errors) >= 1 {
      return errors.New(s.Errors[0].Message)
   }
   *m = s.Data.BonanzaPage.Metadata
   return nil
}

func (m Metadata) Show() string {
   return m.SeriesShortTitle
}

func (m Metadata) Season() int {
   return m.SeasonNumber
}

func (m Metadata) Episode() int {
   return m.EpisodeNumber
}
