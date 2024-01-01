package amc

import (
   "encoding/json"
   "errors"
   "strconv"
   "time"
)

func (v *Video) Unmarshal(c Content) error {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         var s struct {
            Current_Video Video `json:"currentVideo"`
         }
         err := json.Unmarshal(child.Properties, &s)
         if err != nil {
            return nil, err
         }
         *v = s.Current_Video
         return nil
      }
   }
   return errors.New("video-player-ap")
}

func (v Video) Episode() (string, bool) {
   if v.Meta.Episode_Number >= 1 {
      return strconv.Itoa(v.Meta.Episode_Number), true
   }
   return "", false
}

func (Video) Owner() (string, bool) {
   return "", false
}

type Video struct {
   Meta struct {
      Airdate string // 1996-01-01T00:00:00.000Z
      Episode_Number int `json:"episodeNumber"`
      Season int `json:",string"`
      Show_Title string `json:"showTitle"`
   }
   Text struct {
      Title string
   }
}

///////////////////////////////////////////////

func (v Video) Release_Date() (string, bool) {
   return time.Parse(time.RFC3339, v.Meta.Airdate)
}

func (v Video) Season() (int64, error) {
   return v.Meta.Season, nil
}

func (v Video) Series() string {
   return v.Meta.Show_Title
}

func (v Video) Title() string {
   return v.Text.Title
}


