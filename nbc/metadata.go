package nbc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

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

func (m *Metadata) Show() string {
   return m.SeriesShortTitle
}

func (m *Metadata) Season() int {
   return m.SeasonNumber
}

func (m *Metadata) Episode() int {
   return m.EpisodeNumber
}

func (m *Metadata) Year() int {
   return m.AirDate.Year()
}

func (m *Metadata) Title() string {
   if m.MovieShortTitle != "" {
      return m.MovieShortTitle
   }
   return m.SecondaryTitle
}

func (m *Metadata) OnDemand() (*OnDemand, error) {
   req, err := http.NewRequest("", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MpxAccountId, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MpxGuid, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {m.ProgrammingType},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   video := &OnDemand{}
   err = json.NewDecoder(resp.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (m *Metadata) New(guid int) error {
   data, err := func() ([]byte, error) {
      var p page_request
      p.Query = graphql_compact(bonanza_page)
      p.Variables.App = "nbc"
      p.Variables.Name = strconv.Itoa(guid)
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
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         BonanzaPage struct {
            Metadata Metadata
         }
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   if v := value.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   *m = value.Data.BonanzaPage.Metadata
   return nil
}
