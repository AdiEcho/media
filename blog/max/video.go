package max

import (
   "encoding/json"
   "net/http"
   "strings"
   "time"
)

func (d default_token) video(show string) (*active_video, error) {
   req, err := http.NewRequest(
      "", "https://default.any-amer.prd.api.discomax.com", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/content/videos/")
      b.WriteString(show)
      b.WriteString("/activeVideoForShow")
      return b.String()
   }()
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var video struct {
      Data active_video
   }
   err = json.NewDecoder(resp.Body).Decode(&video)
   if err != nil {
      return nil, err
   }
   return &video.Data, nil
}

func (active_video) Show() string {
   return ""
}

func (active_video) Season() int {
   return 0
}

func (active_video) Episode() int {
   return 0
}

func (a active_video) Title() string {
   return a.Attributes.Name
}

type active_video struct {
   Attributes struct {
      AirDate time.Time
      Name string
   }
   Relationships struct {
      Edit struct {
         Data struct {
            Id string
         }
      }
   }
}

func (a active_video) Year() int {
   return a.Attributes.AirDate.Year()
}
