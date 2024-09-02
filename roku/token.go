package roku

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
)

type HomeScreen struct {
   EpisodeNumber int `json:",string"`
   ReleaseDate time.Time // 2007-01-01T000000Z
   SeasonNumber int `json:",string"`
   Series *struct {
      Title string
   }
   Title string
}

func (n Namer) Year() int {
   return n.Home.ReleaseDate.Year()
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
   Home HomeScreen
}

func (n Namer) Show() string {
   if v := n.Home.Series; v != nil {
      return v.Title
   }
   return ""
}

func (n Namer) Season() int {
   return n.Home.SeasonNumber
}

func (n Namer) Episode() int {
   return n.Home.EpisodeNumber
}

func (n Namer) Title() string {
   return n.Home.Title
}

type Playback struct {
   Drm struct {
      Widevine struct {
         LicenseServer string
      }
   }
   Url string
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.Drm.Widevine.LicenseServer, true
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

const user_agent = "trc-googletv; production; 0"

func pointer[T any](value *T) *T {
   return new(T)
}

type AccountToken struct {
   Data []byte
   V    *struct {
      Token string
   }
}

func (a *AccountToken) Unmarshal() error {
   a.V = pointer(a.V)
   return json.Unmarshal(a.Data, a.V)
}
