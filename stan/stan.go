package stan

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (p *LegacyProgram) New(id int64) error {
   address := func() string {
      b := []byte("https://api.stan.com.au/programs/v1/legacy/programs/")
      b = strconv.AppendInt(b, id, 10)
      b = append(b, ".json"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(p)
}

type ActivationCode struct {
   Data []byte
   V struct {
      Code string
      URL string
   }
}

func (a *ActivationCode) New() error {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/activation-codes/", url.Values{
         "generate": {"true"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a ActivationCode) String() string {
   var b strings.Builder
   b.WriteString("Stan.\n")
   b.WriteString("Log in with code\n")
   b.WriteString("1. Visit stan.com.au/activate\n")
   b.WriteString("2. Enter the code:\n")
   b.WriteString(a.V.Code)
   return b.String()
}

func (a ActivationCode) Token() (*WebToken, error) {
   res, err := http.Get(a.V.URL)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var web WebToken
   web.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &web, nil
}

func (a *ActivationCode) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

type AppSession struct {
   JwToken string
}

type WebToken struct {
   Data []byte
   V struct {
      JwToken string
      ProfileId string
   }
}

func (w WebToken) Session() (*AppSession, error) {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/sessions/mobile/app", url.Values{
         "jwToken": {w.V.JwToken},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   session := new(AppSession)
   if err := json.NewDecoder(res.Body).Decode(session); err != nil {
      return nil, err
   }
   return session, nil
}

func (w *WebToken) Unmarshal() error {
   return json.Unmarshal(w.Data, &w.V)
}

type LegacyProgram struct {
   ReleaseYear int
   SeriesTitle string
   Title string
   TvSeasonEpisodeNumber int
   TvSeasonNumber int
}

type Namer struct {
   P LegacyProgram
}

func (n Namer) Episode() int {
   return n.P.TvSeasonEpisodeNumber
}

func (n Namer) Show() string {
   return n.P.SeriesTitle
}

func (n Namer) Season() int {
   return n.P.TvSeasonNumber
}

func (n Namer) Title() string {
   return n.P.Title
}

func (n Namer) Year() int {
   return n.P.ReleaseYear
}
