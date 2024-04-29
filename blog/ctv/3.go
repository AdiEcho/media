package ctv

import (
   "encoding/json"
   "net/http"
   "strconv"
   "time"
)

func (m *media_content) unmarshal(text []byte) error {
   return json.Unmarshal(text, m)
}

func (m media_content) marshal() ([]byte, error) {
   return json.Marshal(m)
}

type media_content struct {
   A *axis_content
   S struct {
      BroadcastDate date
      ContentPackages []struct {
         ID int
      }
      Episode int
      Media struct {
         Name string
         Type string
      }
      Name string
      Season struct {
         Number int
      }
   }
}

func (d *date) UnmarshalText(text []byte) error {
   var err error
   d.T, err = time.Parse(time.DateOnly, string(text))
   if err != nil {
      return err
   }
   return nil
}

func (a axis_content) media() (*media_content, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/")
      b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, a.AxisId, 10)
      b = append(b, "?$include=[ContentPackages,Media,Season]"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   media := media_content{A: &a}
   err = json.NewDecoder(res.Body).Decode(&media.S)
   if err != nil {
      return nil, err
   }
   return &media, nil
}

func (m media_content) Episode() int {
   return m.S.Episode
}

func (m media_content) Season() int {
   return m.S.Season.Number
}

func (m media_content) Show() string {
   if v := m.S.Media; v.Type == "series" {
      return v.Name
   }
   return ""
}

func (m media_content) Title() string {
   if m.S.Media.Type == "movie" {
      return m.S.Name[:len(m.S.Name)-len(" (2024)")]
   }
   return m.S.Name
}

func (m media_content) Year() int {
   return m.S.BroadcastDate.T.Year()
}

type date struct {
   T time.Time
}
