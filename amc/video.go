package amc

import (
   "encoding/json"
   "errors"
   "strconv"
   "strings"
)

func (c ContentCompiler) Video() (*CurrentVideo, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         var s struct {
            CurrentVideo CurrentVideo
         }
         err := json.Unmarshal(child.Properties, &s)
         if err != nil {
            return nil, err
         }
         return &s.CurrentVideo, nil
      }
   }
   return nil, errors.New("video-player-ap")
}

type CurrentVideo struct {
   Meta struct {
      Airdate string // 1996-01-01T00:00:00.000Z
      EpisodeNumber int
      Season int `json:",string"`
      ShowTitle string
   }
   Text struct {
      Title string
   }
}

func (c CurrentVideo) Show() string {
   return c.Meta.ShowTitle
}

func (c CurrentVideo) Season() int {
   return c.Meta.Season
}

func (c CurrentVideo) Episode() int {
   return c.Meta.EpisodeNumber
}

func (c CurrentVideo) Title() string {
   return c.Text.Title
}

func (c CurrentVideo) Year() int {
   if v, _, ok := strings.Cut(c.Meta.Airdate, "-"); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}
