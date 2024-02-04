package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (at App_Token) Item(content_id string) (*Item, error) {
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
      Item_List []Item `json:"itemList"`
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   if len(video.Item_List) == 0 {
      return nil, errors.New("itemList length is zero")
   }
   return &video.Item_List[0], nil
}

type Item struct {
   Air_Date_ISO string `json:"_airDateISO"`
   Label string
   Media_Type string `json:"mediaType"`
   Series_Title string `json:"seriesTitle"`
   // these can be empty string, so we cannot use these:
   // int `json:",string"`
   // json.Number
   Episode_Num string `json:"episodeNum"`
   Season_Num string `json:"seasonNum"`
}

func (Item) Owner() (string, bool) {
   return "", false
}

func (i Item) Show() (string, bool) {
   if i.Media_Type == "Full Episode" {
      return i.Series_Title, true
   }
   return "", false
}

func (i Item) Season() (string, bool) {
   return i.Season_Num, i.Season_Num != ""
}

func (i Item) Episode() (string, bool) {
   return i.Episode_Num, i.Episode_Num != ""
}

func (i Item) Title() (string, bool) {
   return i.Label, true
}

func (i Item) Year() (string, bool) {
   if i.Media_Type == "Movie" {
      year, _, _ := strings.Cut(i.Air_Date_ISO, "-")
      return year, true
   }
   return "", false
}
