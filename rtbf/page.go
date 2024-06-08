package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

// its just not available from what I can tell
func (auvio_page) Year() int {
   return 0
}

func new_page(path string) (*auvio_page, error) {
   res, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      Data auvio_page
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data, nil
}

type auvio_page struct {
   Content struct {
      Title title
      Subtitle subtitle
   }
}

func (a auvio_page) Season() int {
   return a.Content.Title.season
}

type title struct {
   season int
   title string
}

type subtitle struct {
   episode int
   subtitle string
}

// json.data.content.title = "Grantchester S01";
// json.data.content.title = "I care a lot";
func (t *title) UnmarshalText(text []byte) error {
   t.title = string(text)
   if before, after, ok := strings.Cut(t.title, " S"); ok {
      if season, err := strconv.Atoi(after); err == nil {
         t.title = before
         t.season = season
      }
   }
   return nil
}

// json.data.content.subtitle = "06 - Les ombres de la guerre";
// json.data.content.subtitle = "Avec Rosamund Pike";
func (s *subtitle) UnmarshalText(text []byte) error {
   s.subtitle = string(text)
   if before, after, ok := strings.Cut(s.subtitle, " - "); ok {
      if episode, err := strconv.Atoi(before); err == nil {
         s.episode = episode
         s.subtitle = after
      }
   }
   return nil
}

func (a auvio_page) Episode() int {
   return a.Content.Subtitle.episode
}

func (a auvio_page) Show() string {
   if v := a.Content.Title; v.season >= 1 {
      return v.title
   }
   return ""
}

func (a auvio_page) Title() string {
   if v := a.Content.Subtitle; v.episode >= 1 {
      return v.subtitle
   }
   return a.Content.Title.title
}
