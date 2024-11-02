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

func (Authenticate) Marshal(code *LinkCode) ([]byte, error) {
   data, err := json.Marshal(map[string]string{"auth_token": code.AuthToken})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.mubi.com/v3/authenticate", bytes.NewReader(data),
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
   return io.ReadAll(resp.Body)
}

func (a *Authenticate) RequestHeader() (http.Header, error) {
   text, err := json.Marshal(map[string]any{
      "merchant": "mubi",
      "sessionId": a.Token,
      "userId": a.User.Id,
   })
   if err != nil {
      return nil, err
   }
   head := http.Header{}
   head.Set("dt-custom-data", base64.StdEncoding.EncodeToString(text))
   return head, nil
}

// Mubi do this sneaky thing. you cannot download a video unless you have told
// the API that you are watching it. so you have to call
// `/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
// already viewed the video on the website that counts, but if you only use the
// tool it will error
func (a *Authenticate) Viewing(film *FilmResponse) error {
   req, err := http.NewRequest("POST", "https://api.mubi.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = func() string {
      b := []byte("/v3/films/")
      b = strconv.AppendInt(b, film.Id, 10)
      b = append(b, "/viewing"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Token},
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
   return nil
}

func (*Authenticate) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

// final slash is needed
func (*Authenticate) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (*Authenticate) UnwrapResponse(b []byte) ([]byte, error) {
   var data struct {
      License []byte
   }
   err := json.Unmarshal(b, &data)
   if err != nil {
      return nil, err
   }
   return data.License, nil
}

type Authenticate struct {
   Token string
   User struct {
      Id int
   }
}

func (a *Authenticate) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}
