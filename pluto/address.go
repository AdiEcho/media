package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (a Address) Video(forward string) (*OnDemand, error) {
   req, err := http.NewRequest("", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName":           {"web"},
      "appVersion":        {"9"},
      "clientID":          {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities":   {"widevine:L3"},
      "seriesIDs":         {a[0]},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Vod []OnDemand
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   demand := value.Vod[0]
   if demand.Slug != a[0] {
      if demand.Id != a[0] {
         return nil, errors.New(demand.Slug)
      }
   }
   for _, season := range demand.Seasons {
      season.show = &demand
      for _, episode := range season.Episodes {
         episode.season = season
         if episode.Episode == a[1] {
            return episode, nil
         }
         if episode.Slug == a[1] {
            return episode, nil
         }
      }
   }
   return &demand, nil
}

func (a Address) String() string {
   var b strings.Builder
   if a[0] != "" {
      if a[1] != "" {
         b.WriteString("series/")
         b.WriteString(a[0])
         b.WriteString("/episode/")
         b.WriteString(a[1])
      } else {
         b.WriteString("movies/")
         b.WriteString(a[0])
      }
   }
   return b.String()
}

type Address [2]string

func (a *Address) Set(text string) error {
   for {
      var (
         key string
         ok  bool
      )
      key, text, ok = strings.Cut(text, "/")
      if !ok {
         return nil
      }
      switch key {
      case "movies":
         (*a)[0] = text
      case "series":
         (*a)[0], text, ok = strings.Cut(text, "/")
         if !ok {
            return errors.New("episode")
         }
      case "episode":
         (*a)[1] = text
      }
   }
}
