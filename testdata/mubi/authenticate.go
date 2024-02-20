package mubi

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

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

func (c linkCode) authenticate() (*authenticate, error) {
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

type authenticate struct {
   s struct {
      Token string
      User struct {
         ID int
      }
   }
   Raw []byte
}
