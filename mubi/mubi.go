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

func (film_response) Owner() (string, bool) {
   return "", false
}

func (film_response) Show() (string, bool) {
   return "", false
}

func (film_response) Season() (string, bool) {
   return "", false
}

func (film_response) Episode() (string, bool) {
   return "", false
}

func (f film_response) Title() (string, bool) {
   return f.s.Title, true
}

func (f film_response) Year() (string, bool) {
   return strconv.Itoa(f.s.Year), true
}

type secure_url struct {
   URL string
}

func (c link_code) authenticate() (*authenticate, error) {
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
   var auth authenticate
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

func (c *link_code) New() error {
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

func (c link_code) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.s.Link_Code)
   return b.String()
}

type link_code struct {
   Raw []byte
   s struct {
      Auth_Token string
      Link_Code string
   }
}

func (c *link_code) unmarshal() error {
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

type film_response struct {
   s struct {
      Title string
      Year int
   }
}

type WebAddress struct {
   s string
}

func (w WebAddress) film() (*film_response, error) {
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
   var film film_response
   if err := json.NewDecoder(res.Body).Decode(&film.s); err != nil {
      return nil, err
   }
   return &film, nil
}
