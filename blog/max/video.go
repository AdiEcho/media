package max

import (
   "encoding/json"
   "net/http"
   "strings"
)

type active_video struct {
   Data struct {
      Relationships struct {
         Edit struct {
            Data struct {
               Id string
            }
         }
      }
   }
}

func (d default_token) video() (*active_video, error) {
   req, err := http.NewRequest(
      "", "https://default.any-amer.prd.api.discomax.com", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/content/videos/")
      b.WriteString("127b00c5-0131-4bac-b2d1-40762deefe09")
      b.WriteString("/activeVideoForShow")
      return b.String()
   }()
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   video := new(active_video)
   err = json.NewDecoder(resp.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}
