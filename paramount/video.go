package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
   "time"
)

type number int

func (n *number) UnmarshalText(text []byte) error {
   if len(text) >= 1 {
      i, err := strconv.Atoi(string(text))
      if err != nil {
         return err
      }
      *n = number(i)
   }
   return nil
}

func (v VideoItem) Season() int {
   return int(v.SeasonNum)
}

func (v VideoItem) Episode() int {
   return int(v.EpisodeNum)
}

func (v VideoItem) Title() string {
   return v.Label
}

func (v VideoItem) Year() int {
   return v.AirDateIso.Year()
}

func (v VideoItem) Show() string {
   if v.MediaType == "Full Episode" {
      return v.SeriesTitle
   }
   return ""
}

type VideoItem struct {
   SeriesTitle string
   SeasonNum number
   EpisodeNum number
   Label string
   AirDateIso time.Time `json:"_airDateISO"`
   MediaType string
}

// must use app token and IP address for correct location
func (at AppToken) Item(content_id string) (*VideoItem, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v2.0/androidphone/video/cid/")
      b.WriteString(content_id)
      b.WriteString(".json")
      return b.String()
   }()
   req.URL.RawQuery = at.v.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var video struct {
      Error string
      ItemList []VideoItem
   }
   err = json.NewDecoder(resp.Body).Decode(&video)
   if err != nil {
      return nil, err
   }
   if video.Error != "" {
      return nil, errors.New(video.Error)
   }
   if len(video.ItemList) == 0 {
      return nil, errors.New("len(itemList) == 0")
   }
   return &video.ItemList[0], nil
}
