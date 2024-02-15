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

func (CurrentVideo) Owner() (string, bool) {
   return "", false
}

func (c CurrentVideo) Season() (string, bool) {
   if c.Meta.Season >= 1 {
      return strconv.Itoa(c.Meta.Season), true
   }
   return "", false
}

func (c CurrentVideo) Episode() (string, bool) {
   if c.Meta.EpisodeNumber >= 1 {
      return strconv.Itoa(c.Meta.EpisodeNumber), true
   }
   return "", false
}

func (c CurrentVideo) Title() (string, bool) {
   return c.Text.Title, true
}

func (c CurrentVideo) Show() (string, bool) {
   return c.Meta.ShowTitle, c.Meta.ShowTitle != ""
}

func (c CurrentVideo) Year() (string, bool) {
   if c.Meta.ShowTitle != "" {
      return "", false
   }
   year, _, _ := strings.Cut(c.Meta.Airdate, "-")
   return year, true
}
