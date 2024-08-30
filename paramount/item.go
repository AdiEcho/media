package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
   "time"
)

type VideoItem struct {
   SeriesTitle string
   SeasonNum Number
   EpisodeNum Number
   Label string
   AirDateIso time.Time `json:"_airDateISO"`
   MediaType string
}

func (v *VideoItem) Season() int {
   return int(v.SeasonNum)
}

func (v *VideoItem) Episode() int {
   return int(v.EpisodeNum)
}

func (v *VideoItem) Title() string {
   return v.Label
}

func (v *VideoItem) Year() int {
   return v.AirDateIso.Year()
}

///

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
   var value struct {
      Error string
      ItemList []VideoItem
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if v := value.Error; v != "" {
      return nil, errors.New(v)
   }
   return &value.ItemList[0], nil
}

func (v *VideoItem) Json(text []byte) error {
   return json.Unmarshal(text, v)
}

func (v VideoItems) Item() (*VideoItem, bool) {
   if len(v) >= 1 {
      return &v[0], true
   }
   return nil, false
}

func (v VideoItem) Show() string {
   if v.MediaType == "Full Episode" {
      return v.SeriesTitle
   }
   return ""
}

type VideoItems []VideoItem
