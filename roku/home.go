package roku

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
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
   return json.NewDecoder(res.Body).Decode(&h.V)
}

func (h HomeScreen) DASH() (*MediaVideo, bool) {
   for _, option := range h.V.ViewOptions {
      for _, video := range option.Media.Videos {
         if video.VideoType == "DASH" {
            return &video, true
         }
      }
   }
   return nil, false
}

// we have to embed to prevent clobbering Namer.Title
type HomeScreen struct {
   V struct {
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

func (h HomeScreen) Show() string {
   return h.V.Series.Title
}

func (h HomeScreen) Season() int {
   return h.V.SeasonNumber
}

func (h HomeScreen) Episode() int {
   return h.V.EpisodeNumber
}

func (h HomeScreen) Title() string {
   return h.V.Title
}

func (h HomeScreen) Year() int {
   if v, _, ok := strings.Cut(h.V.ReleaseDate, "-"); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}
