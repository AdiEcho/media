package roku

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (h *HomeScreen) New(id string) error {
   // outside of US this redirects to
   // therokuchannel.roku.com/enguard
   // so disable redirect
   client := http.Client{
      CheckRedirect: func(*http.Request, []*http.Request) error {
         return http.ErrUseLastResponse
      },
   }
   // this has got to be the stupidest fucking URL construction I have ever
   // seen, but its what the server requires
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
   res, err := client.Get(b.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(res.Body).Decode(h)
}

func (h HomeScreen) DASH() (*MediaVideo, bool) {
   for _, option := range h.ViewOptions {
      for _, video := range option.Media.Videos {
         if video.VideoType == "DASH" {
            return &video, true
         }
      }
   }
   return nil, false
}

type HomeScreen struct {
   EpisodeNumber int `json:",string"`
   ReleaseDate string // 2007-01-01T000000Z
   SeasonNumber int `json:",string"`
   Series *struct {
      Title string
   }
   Title string
   ViewOptions []struct {
      Media struct {
         Videos []MediaVideo
      }
   }
}

type Name struct {
   H HomeScreen
}

func (n Name) Show() string {
   if v := n.H.Series; v != nil {
      return v.Title
   }
   return ""
}

func (n Name) Season() int {
   return n.H.SeasonNumber
}

func (n Name) Episode() int {
   return n.H.EpisodeNumber
}

func (n Name) Title() string {
   return n.H.Title
}

func (n Name) Year() int {
   if v, _, ok := strings.Cut(n.H.ReleaseDate, "-"); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}
