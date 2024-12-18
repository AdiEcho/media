package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (p *Playback) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      p.Drm.Widevine.LicenseServer, "application/x-protobuf",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Playback struct {
   Drm struct {
      Widevine struct {
         LicenseServer string
      }
   }
   Url string
}

func (AccountToken) Marshal(
   auth *AccountAuth, code *AccountCode,
) ([]byte, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + code.Code
   req.Header = http.Header{
      "user-agent":           {user_agent},
      "x-roku-content-token": {auth.AuthToken},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

const user_agent = "trc-googletv; production; 0"

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

type AccountToken struct {
   Token string
}

func (a *AccountToken) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (n *Namer) Title() string {
   return n.Home.Title
}

func (n *Namer) Show() string {
   if v := n.Home.Series; v != nil {
      return v.Title
   }
   return ""
}

type Namer struct {
   Home HomeScreen
}

type HomeScreen struct {
   EpisodeNumber int       `json:",string"`
   ReleaseDate   time.Time // 2007-01-01T000000Z
   SeasonNumber  int       `json:",string"`
   Series        *struct {
      Title string
   }
   Title string
}

func (n *Namer) Season() int {
   return n.Home.SeasonNumber
}

func (n *Namer) Episode() int {
   return n.Home.EpisodeNumber
}

func (n *Namer) Year() int {
   return n.Home.ReleaseDate.Year()
}
