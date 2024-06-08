package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

func (a AuvioPage) asset_id() string {
   if v := a.Content.AssetId; v != "" {
      return v
   }
   return a.Content.Media.AssetId
}

type AuvioPage struct {
   Content struct {
      AssetId  string
      Media struct {
         AssetId string
      }
      Subtitle Subtitle
      Title    Title
   }
}

type Title struct {
   Season int
   Title  string
}

type Subtitle struct {
   Episode  int
   Subtitle string
}

func NewPage(path string) (*AuvioPage, error) {
   res, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      Data AuvioPage
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data, nil
}

func (a AuvioPage) Season() int {
   return a.Content.Title.Season
}

// json.data.content.title = "Grantchester S01";
// json.data.content.title = "I care a lot";
func (t *Title) UnmarshalText(text []byte) error {
   t.Title = string(text)
   if before, after, ok := strings.Cut(t.Title, " S"); ok {
      if season, err := strconv.Atoi(after); err == nil {
         t.Title = before
         t.Season = season
      }
   }
   return nil
}

// json.data.content.subtitle = "06 - Les ombres de la guerre";
// json.data.content.subtitle = "Avec Rosamund Pike";
func (s *Subtitle) UnmarshalText(text []byte) error {
   s.Subtitle = string(text)
   if before, after, ok := strings.Cut(s.Subtitle, " - "); ok {
      if episode, err := strconv.Atoi(before); err == nil {
         s.Episode = episode
         s.Subtitle = after
      }
   }
   return nil
}

func (a AuvioPage) Episode() int {
   return a.Content.Subtitle.Episode
}

func (a AuvioPage) Show() string {
   if v := a.Content.Title; v.Season >= 1 {
      return v.Title
   }
   return ""
}

func (a AuvioPage) Title() string {
   if v := a.Content.Subtitle; v.Episode >= 1 {
      return v.Subtitle
   }
   return a.Content.Title.Title
}

// its just not available from what I can tell
func (AuvioPage) Year() int {
   return 0
}
