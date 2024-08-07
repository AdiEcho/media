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
   Id string
   Url string
}

func (t TextTrack) String() string {
   var b strings.Builder
   b.WriteString("id = ")
   b.WriteString(t.Id)
   return b.String()
}

type SecureUrl struct {
   Data []byte
   V struct {
      TextTrackUrls []TextTrack `json:"text_track_urls"`
      Url string
   }
}

func (s *SecureUrl) Unmarshal() error {
   return json.Unmarshal(s.Data, &s.V)
}

var ClientCountry = "US"

// "android" requires headers:
// client-device-identifier
// client-version
const client = "web"

func (Authenticate) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

// final slash is needed
func (Authenticate) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (Authenticate) UnwrapResponse(b []byte) ([]byte, error) {
   var data struct {
      License []byte
   }
   err := json.Unmarshal(b, &data)
   if err != nil {
      return nil, err
   }
   return data.License, nil
}

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

func (c *LinkCode) New() error {
   req, err := http.NewRequest("", "https://api.mubi.com/v3/link_code", nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   c.Data, err = io.ReadAll(resp.Body)
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

func (a Authenticate) RequestHeader() (http.Header, error) {
   value := map[string]any{
      "merchant": "mubi",
      "sessionId": a.V.Token,
      "userId": a.V.User.Id,
   }
   text, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   head := http.Header{}
   head.Set("dt-custom-data", base64.StdEncoding.EncodeToString(text))
   return head, nil
}

type Authenticate struct {
   Data []byte
   V struct {
      Token string
      User struct {
         Id int
      }
   }
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
      "client": {client},
      "client-country": {ClientCountry},
      "content-type": {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &Authenticate{Data: data}, nil
}
