package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (at AppToken) Item(content_id string) (*Item, error) {
   req, err := http.NewRequest("GET", "https://www.paramountplus.com", nil)
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
   // this needs to be encoded
   req.URL.RawQuery = "at=" + url.QueryEscape(at.value)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var video struct {
      ItemList []Item
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   if len(video.ItemList) == 0 {
      return nil, errors.New("itemList length is zero")
   }
   return &video.ItemList[0], nil
}

type Item struct {
   AirDateIso string `json:"_airDateISO"`
   Label string
   MediaType string
   SeriesTitle string
   // these can be empty string, so we cannot use these:
   // int `json:",string"`
   // json.Number
   EpisodeNum string
   SeasonNum string
}

func (Item) Owner() (string, bool) {
   return "", false
}

func (i Item) Show() (string, bool) {
   if i.MediaType == "Full Episode" {
      return i.SeriesTitle, true
   }
   return "", false
}

func (i Item) Season() (string, bool) {
   return i.SeasonNum, i.SeasonNum != ""
}

func (i Item) Episode() (string, bool) {
   return i.EpisodeNum, i.EpisodeNum != ""
}

func (i Item) Title() (string, bool) {
   return i.Label, true
}

func (i Item) Year() (string, bool) {
   if i.MediaType == "Movie" {
      year, _, _ := strings.Cut(i.AirDateIso, "-")
      return year, true
   }
   return "", false
}
