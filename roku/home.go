package roku

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (n Namer) Year() int {
   return n.H.ReleaseDate.Year()
}

type HomeScreen struct {
   EpisodeNumber int `json:",string"`
   ReleaseDate time.Time // 2007-01-01T000000Z
   SeasonNumber int `json:",string"`
   Series *struct {
      Title string
   }
   Title string
}

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
         b.WriteString("title")
         return url.PathEscape(b.String())
      }())
      return url.PathEscape(b.String())
   }())
   resp, err := client.Get(b.String())
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(resp.Body).Decode(h)
}

type Namer struct {
   H HomeScreen
}

func (n Namer) Show() string {
   if v := n.H.Series; v != nil {
      return v.Title
   }
   return ""
}

func (n Namer) Season() int {
   return n.H.SeasonNumber
}

func (n Namer) Episode() int {
   return n.H.EpisodeNumber
}

func (n Namer) Title() string {
   return n.H.Title
}
