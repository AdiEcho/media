package ctv

import (
   "encoding/json"
   "net/http"
   "strconv"
   "time"
)

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
   media := new(media_content)
   err = json.NewDecoder(res.Body).Decode(media)
   if err != nil {
      return nil, err
   }
   return media, nil
}

type date struct {
   t time.Time
}

func (d *date) UnmarshalText(text []byte) error {
   var err error
   d.t, err = time.Parse(time.DateOnly, string(text))
   if err != nil {
      return err
   }
   return nil
}

type media_content struct {
   BroadcastDate date
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

type namer struct {
   m *media_content
}

func (n namer) Episode() int {
   return n.m.Episode
}

func (n namer) Season() int {
   return n.m.Season.Number
}

func (n namer) Show() string {
   if v := n.m.Media; v.Type == "series" {
      return v.Name
   }
   return ""
}

func (n namer) Title() string {
   if n.m.Media.Type == "movie" {
      return n.m.Name[:len(n.m.Name)-len(" (2024)")]
   }
   return n.m.Name
}

func (n namer) Year() int {
   return n.m.BroadcastDate.t.Year()
}
