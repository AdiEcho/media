package mubi

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type TextTrack struct {
   ID string
   URL string
}

func (t TextTrack) String() string {
   var b strings.Builder
   b.WriteString("id = ")
   b.WriteString(t.ID)
   return b.String()
}

type SecureUrl struct {
   Data []byte
   V struct {
      TextTrackUrls []TextTrack `json:"text_track_urls"`
      URL string
   }
}

func (s *SecureUrl) Unmarshal() error {
   return json.Unmarshal(s.Data, &s.V)
}

var ClientCountry = "US"

// "android" requires headers:
// Client-Device-Identifier
// Client-Version
const client = "web"

type Authenticate struct {
   Data []byte
   V struct {
      Token string
      User struct {
         ID int
      }
   }
}

func (Authenticate) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (a Authenticate) RequestHeader() (http.Header, error) {
   value := map[string]any{
      "merchant": "mubi",
      "sessionId": a.V.Token,
      "userId": a.V.User.ID,
   }
   text, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   head := make(http.Header)
   head.Set("Dt-Custom-Data", base64.StdEncoding.EncodeToString(text))
   return head, nil
}

// final slash is needed
func (Authenticate) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (Authenticate) UnwrapResponse(b []byte) ([]byte, error) {
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
   return json.Unmarshal(a.Data, &a.V)
}

func (c LinkCode) Authenticate() (*Authenticate, error) {
   body, err := json.Marshal(map[string]string{"auth_token": c.V.AuthToken})
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
   b.WriteString(c.V.LinkCode)
   return b.String()
}

func (c *LinkCode) Unmarshal() error {
   return json.Unmarshal(c.Data, &c.V)
}

type LinkCode struct {
   Data []byte
   V struct {
      AuthToken string `json:"auth_token"`
      LinkCode string `json:"link_code"`
   }
}

type Address struct {
   s string
}

func (a *Address) Set(text string) error {
   var ok bool
   _, a.s, ok = strings.Cut(text, "/films/")
   if !ok {
      return errors.New("/films/")
   }
   return nil
}

func (a Address) String() string {
   return a.s
}
