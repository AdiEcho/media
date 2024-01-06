package amc

import (
   "encoding/json"
   "errors"
   "strconv"
   "strings"
)

func (c Content) Video() (*Video, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         var s struct {
            Current_Video Video `json:"currentVideo"`
         }
         err := json.Unmarshal(child.Properties, &s)
         if err != nil {
            return nil, err
         }
         return &s.Current_Video, nil
      }
   }
   return nil, errors.New("video-player-ap")
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

func (Video) Owner() (string, bool) {
   return "", false
}

func (v Video) Season() (string, bool) {
   if v.Meta.Season >= 1 {
      return strconv.Itoa(v.Meta.Season), true
   }
   return "", false
}

func (v Video) Episode() (string, bool) {
   if v.Meta.Episode_Number >= 1 {
      return strconv.Itoa(v.Meta.Episode_Number), true
   }
   return "", false
}

func (v Video) Title() (string, bool) {
   return v.Text.Title, true
}

func (v Video) Show() (string, bool) {
   return v.Meta.Show_Title, v.Meta.Show_Title != ""
}

func (v Video) Year() (string, bool) {
   if v.Meta.Show_Title != "" {
      return "", false
   }
   year, _, _ := strings.Cut(v.Meta.Airdate, "-")
   return year, true
}
