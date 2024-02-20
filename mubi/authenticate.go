package mubi

import (
   "encoding/base64"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

func (a authenticate) secure(film int64) (*secure_url, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, film, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }
   req, err := http.NewRequest("GET", address(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.s.Token},
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   secure := new(secure_url)
   if err := json.NewDecoder(res.Body).Decode(secure); err != nil {
      return nil, err
   }
   return secure, nil
}

// final slash is needed
func (authenticate) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (authenticate) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (authenticate) ResponseBody(b []byte) ([]byte, error) {
   var s struct {
      License []byte
   }
   err := json.Unmarshal(b, &s)
   if err != nil {
      return nil, err
   }
   return s.License, nil
}

func (a authenticate) RequestHeader() (http.Header, error) {
   value := map[string]any{
      "merchant": "mubi",
      "sessionId": a.s.Token,
      "userId": a.s.User.ID,
   }
   text, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   head := make(http.Header)
   head.Set("Dt-Custom-Data", base64.StdEncoding.EncodeToString(text))
   return head, nil
}

func (a *authenticate) unmarshal() error {
   return json.Unmarshal(a.Raw, &a.s)
}

type authenticate struct {
   s struct {
      Token string
      User struct {
         ID int
      }
   }
   Raw []byte
}
