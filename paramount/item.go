package paramount

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
   "time"
)

func (v *VideoItem) asset_type() string {
   if v.MediaType == "Movie" {
      return "DASH_CENC_PRECON"
   }
   return "DASH_CENC"
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

func (v *VideoItem) Show() string {
   if v.MediaType == "Full Episode" {
      return v.SeriesTitle
   }
   return ""
}

type VideoItem struct {
   AirDateIso time.Time `json:"_airDateISO"`
   CmsAccountId string
   ContentId string
   EpisodeNum Number
   Label string
   MediaType string
   SeasonNum Number
   SeriesTitle string
}

func (v *VideoItem) Unmarshal(data []byte) error {
   var value struct {
      Error string
      ItemList []VideoItem
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   if value.Error != "" {
      return errors.New(value.Error)
   }
   if len(value.ItemList) == 0 {
      return errors.New(`"itemList":[]`)
   }
   *v = value.ItemList[0]
   return nil
}

// must use app token and IP address for correct location
func (a *AppToken) Item(cid string, data *[]byte) (*VideoItem, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v2.0/androidphone/video/cid/")
      b.WriteString(cid)
      b.WriteString(".json")
      return b.String()
   }()
   req.URL.RawQuery = a.Values.Encode()
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
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if data != nil {
      *data = body
      return nil, nil
   }
   var item VideoItem
   err = item.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   return &item, nil
}
