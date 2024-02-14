package roku

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

// this has got to be the stupidest fucking URL construction I have ever seen,
// but its what the server requires
func (h *HomeScreen) New(id string) error {
   var b strings.Builder
   b.WriteString("https://therokuchannel.roku.com/api/v2/homescreen/content/")
   b.WriteString(func() string {
      var b strings.Builder
      b.WriteString("https://content.sr.roku.com/content/v1/roku-trc/")
      b.WriteString(id)
      b.WriteByte('?')
      b.WriteString("expand=series&")
      b.WriteString("include=")
      b.WriteString(func() string {
         var b strings.Builder
         b.WriteString("episodeNumber,")
         b.WriteString("releaseDate,")
         b.WriteString("seasonNumber,")
         b.WriteString("series.title,")
         b.WriteString("title,")
         b.WriteString("viewOptions")
         return url.PathEscape(b.String())
      }())
      return url.PathEscape(b.String())
   }())
   res, err := http.Get(b.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(&h.s)
}

// we have to embed to prevent clobbering Namer.Title
type HomeScreen struct {
   s struct {
      Series *struct {
         Title string
      }
      Title string
      ReleaseDate string // 2007-01-01T000000Z
      ViewOptions []struct {
         Media struct {
            Videos []MediaVideo
         }
      }
      SeasonNumber string
      EpisodeNumber string
   }
}

func (h HomeScreen) Show() (string, bool) {
   if h.s.Series != nil {
      return h.s.Series.Title, true
   }
   return "", false
}

func (h HomeScreen) Year() (string, bool) {
   if h.s.Series != nil {
      return "", false
   }
   year, _, _ := strings.Cut(h.s.ReleaseDate, "-")
   return year, true
}

func (HomeScreen) Owner() (string, bool) {
   return "", false
}

func (h HomeScreen) Title() (string, bool) {
   return h.s.Title, true
}

func (h HomeScreen) Season() (string, bool) {
   if sn := h.s.SeasonNumber; sn != "" {
      return sn, true
   }
   return "", false
}

func (h HomeScreen) Episode() (string, bool) {
   if en := h.s.EpisodeNumber; en != "" {
      return en, true
   }
   return "", false
}

func (h HomeScreen) DASH() (*MediaVideo, bool) {
   for _, option := range h.s.ViewOptions {
      for _, video := range option.Media.Videos {
         if video.VideoType == "DASH" {
            return &video, true
         }
      }
   }
   return nil, false
}
