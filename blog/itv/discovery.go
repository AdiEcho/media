package itv

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
)

const query_discovery = `
{
   titles(filter: {
      legacyId: %q
   }) {
      brand {
         title
      }
      ... on Episode {
         seriesNumber
         episodeNumber
      }
      title
      ... on Film {
         productionYear
      }
   }
}
`

func (d *discovery_title) New(legacy_id string) error {
   req, err := http.NewRequest(
      "", "https://content-inventory.prd.oasvc.itv.com/discovery", nil,
   )
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "query": {fmt.Sprintf(query_discovery, legacy_id)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Titles []discovery_title
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   *d = value.Data.Titles[0]
   return nil
}

type discovery_title struct {
   Brand *struct {
      Title string
   }
   SeriesNumber int
   EpisodeNumber int
   Title string
   ProductionYear int
}
