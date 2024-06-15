package max

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
   "time"
)

type address struct {
   video_id string
   edit_id string
}

type route_include struct {
   Attributes struct {
      AirDate time.Time
      Name string
      EpisodeNumber int
      SeasonNumber int
   }
   Id string
   Relationships *struct {
      Show *struct {
         Data struct {
            Id string
         }
      }
   }
}

func (d default_token) routes(path string) (*default_routes, error) {
   req, err := http.NewRequest(
      "", "https://default.any-amer.prd.api.discomax.com/cms/routes"+path, nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   req.URL.RawQuery = url.Values{
      "include": {"default"},
      // this is not required, but results in a smaller response
      "page[items.size]": {"1"},
   }.Encode()
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

func (a *address) UnmarshalText(text []byte) error {
   split := strings.Split(string(text), "/")
   a.video_id = split[3]
   a.edit_id = split[4]
   return nil
}

type default_routes struct {
   Data struct {
      Attributes struct {
         Url address
      }
   }
   Included []route_include
}

func (d default_routes) video() (*route_include, bool) {
   for _, include := range d.Included {
      if include.Id == d.Data.Attributes.Url.video_id {
         return &include, true
      }
   }
   return nil, false
}

func (d default_routes) Show() string {
   if v, ok := d.video(); ok {
      if v.Attributes.SeasonNumber >= 1 {
         for _, include := range d.Included {
            if include.Id == v.Relationships.Show.Data.Id {
               return include.Attributes.Name
            }
         }
      }
   }
   return ""
}

func (d default_routes) Season() int {
   if v, ok := d.video(); ok {
      return v.Attributes.SeasonNumber
   }
   return 0
}

func (d default_routes) Episode() int {
   if v, ok := d.video(); ok {
      return v.Attributes.EpisodeNumber
   }
   return 0
}

func (d default_routes) Title() string {
   if v, ok := d.video(); ok {
      return v.Attributes.Name
   }
   return ""
}

func (d default_routes) Year() int {
   if v, ok := d.video(); ok {
      return v.Attributes.AirDate.Year()
   }
   return 0
}
