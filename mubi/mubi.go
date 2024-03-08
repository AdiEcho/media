package mubi

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

type Authenticate struct {
   Data []byte
   v struct {
      Token string
      User struct {
         ID int
      }
   }
}

func (a Authenticate) RequestHeader() (http.Header, error) {
   value := map[string]any{
      "merchant": "mubi",
      "sessionId": a.v.Token,
      "userId": a.v.User.ID,
   }
   text, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   head := make(http.Header)
   head.Set("Dt-Custom-Data", base64.StdEncoding.EncodeToString(text))
   return head, nil
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

// Mubi do this sneaky thing. you cannot download a video unless you have told
// the API that you are watching it. so you have to call
// `/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
// already viewed the video on the website that counts, but if you only use the
// tool it will error
func (a Authenticate) Viewing(f *FilmResponse) error {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, f.v.ID, 10)
      b = append(b, "/viewing"...)
      return string(b)
   }
   req, err := http.NewRequest("POST", address(), nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.v.Token},
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
   return nil
}

// final slash is needed
func (Authenticate) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (Authenticate) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (Authenticate) ResponseBody(b []byte) ([]byte, error) {
   var v struct {
      License []byte
   }
   err := json.Unmarshal(b, &v)
   if err != nil {
      return nil, err
   }
   return v.License, nil
}

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.v)
}

func (w WebAddress) Film() (*FilmResponse, error) {
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
   if err := json.NewDecoder(res.Body).Decode(&film.v); err != nil {
      return nil, err
   }
   return &film, nil
}

type FilmResponse struct {
   v struct {
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
   return f.v.Title, true
}

func (f FilmResponse) Year() (string, bool) {
   return strconv.Itoa(f.v.Year), true
}

func (c LinkCode) Authenticate() (*Authenticate, error) {
   body, err := json.Marshal(map[string]string{"auth_token": c.v.Auth_Token})
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
   auth.Data, err = io.ReadAll(res.Body)
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
   c.Data, err = io.ReadAll(res.Body)
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
   b.WriteString(c.v.Link_Code)
   return b.String()
}

type LinkCode struct {
   Data []byte
   v struct {
      Auth_Token string
      Link_Code string
   }
}

func (c *LinkCode) Unmarshal() error {
   return json.Unmarshal(c.Data, &c.v)
}

