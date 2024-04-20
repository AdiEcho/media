package pluto

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "strings"
)

func (w web_address) String() string {
   var b strings.Builder
   b.WriteString("https://pluto.tv/on-demand/")
   if w.episode != "" {
      b.WriteString("series")
   } else {
      b.WriteString("movies")
   }
   b.WriteByte('/')
   b.WriteString(w.series)
   if w.episode != "" {
      b.WriteString("/episode/")
      b.WriteString(w.episode)
   }
   return b.String()
}

func (w *web_address) Set(s string) error {
   for {
      var (
         key string
         ok bool
      )
      key, s, ok = strings.Cut(s, "/")
      if !ok {
         return nil
      }
      switch key {
      case "episode":
         w.episode = s
      case "movies":
         w.series = s
      case "series":
         w.series, s, _ = strings.Cut(s, "/")
      }
   }
}

func (w web_address) video() (*video, error) {
   req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "appName": {"web"},
      "appVersion": {"9"},
      "clientID": {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities": {"widevine:L3"},
      "seriesIDs": {w.series},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var start struct {
      VOD []video
   }
   err = json.NewDecoder(res.Body).Decode(&start)
   if err != nil {
      return nil, err
   }
   demand := start.VOD[0]
   if demand.Slug != w.series {
      if demand.ID != w.series {
         return nil, fmt.Errorf("%+v", demand)
      }
   }
   for _, s := range demand.Seasons {
      s.parent = &demand
      for _, e := range s.Episodes {
         e.parent = s
         if e.Episode == w.episode {
            return e, nil
         }
         if e.Slug == w.episode {
            return e, nil
         }
      }
   }
   return &demand, nil
}

type web_address struct {
   series string
   episode string
}

type video struct {
   Episode string `json:"_id"`
   ID string
   Name string
   Seasons []*season
   Slug string
   parent *season
}

type season struct {
   Episodes []*video
   parent *video
}
