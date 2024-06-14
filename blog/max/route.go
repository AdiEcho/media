package max

import (
   "encoding/json"
   "net/http"
   "time"
)

func (d default_token) routes(path string) (*default_routes, error) {
   req, err := http.NewRequest(
      "", "https://default.any-amer.prd.api.discomax.com/cms/routes"+path, nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "include=default"
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   route := new(default_routes)
   err = json.NewDecoder(resp.Body).Decode(route)
   if err != nil {
      return nil, err
   }
   return route, nil
}

type default_routes struct {
   Included []struct {
      Attributes struct {
         AirDate time.Time
         EpisodeNumber int
         Name string
         SeasonNumber int
         Type string
      }
   }
}
