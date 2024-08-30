package mubi

import (
   "encoding/base64"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

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

// Mubi do this sneaky thing. you cannot download a video unless you have told
// the API that you are watching it. so you have to call
// `/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
// already viewed the video on the website that counts, but if you only use the
// tool it will error
func (a Authenticate) Viewing(film *FilmResponse) error {
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
      "authorization": {"Bearer " + a.V.Token},
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

