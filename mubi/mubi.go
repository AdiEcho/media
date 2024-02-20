package mubi

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

type FilmResponse struct {
   s struct {
      ID int64
      Title string
      Year int
   }
}

func (FilmResponse) Owner() (string, bool) {
   return "", false
}

func (FilmResponse) Show() (string, bool) {
   return "", false
}

func (FilmResponse) Season() (string, bool) {
   return "", false
}

func (FilmResponse) Episode() (string, bool) {
   return "", false
}

func (f FilmResponse) Title() (string, bool) {
   return f.s.Title, true
}

func (f FilmResponse) Year() (string, bool) {
   return strconv.Itoa(f.s.Year), true
}

type SecureUrl struct {
   URL string
}

func (c LinkCode) Authenticate() (*Authenticate, error) {
   body, err := json.Marshal(map[string]string{"auth_token": c.s.Auth_Token})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.mubi.com/v3/authenticate", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Client": {client},
      "Client-Country": {ClientCountry},
      "Content-Type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var auth Authenticate
   auth.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

// "android" requires headers:
// Client-Device-Identifier
// Client-Version
const client = "web"

var ClientCountry = "US"

func (c *LinkCode) New() error {
   req, err := http.NewRequest("GET", "https://api.mubi.com/v3/link_code", nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   c.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (c LinkCode) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.s.Link_Code)
   return b.String()
}

type LinkCode struct {
   Raw []byte
   s struct {
      Auth_Token string
      Link_Code string
   }
}

func (c *LinkCode) Unmarshal() error {
   return json.Unmarshal(c.Raw, &c.s)
}

func (w WebAddress) String() string {
   return w.s
}

func (w *WebAddress) Set(s string) error {
   var ok bool
   _, w.s, ok = strings.Cut(s, "/films/")
   if !ok {
      return errors.New("/films/")
   }
   return nil
}

type WebAddress struct {
   s string
}

func (w WebAddress) film() (*FilmResponse, error) {
   req, err := http.NewRequest(
      "GET", "https://api.mubi.com/v3/films/" + w.s, nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var film FilmResponse
   if err := json.NewDecoder(res.Body).Decode(&film.s); err != nil {
      return nil, err
   }
   return &film, nil
}
