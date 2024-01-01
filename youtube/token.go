package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
   "os"
)

// YouTube on TV
const (
   client_ID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func (d Device_Code) Token() (Raw_Token, error) {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_ID},
         "client_secret": {client_secret},
         "device_code": {d.Device_Code},
         "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

type Raw_Token []byte

type Token struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func (t *Token) Refresh(r Raw_Token) error {
   err := json.Unmarshal(r, t)
   if err != nil {
      return err
   }
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_ID},
         "client_secret": {client_secret},
         "grant_type": {"refresh_token"},
         "refresh_token": {t.Refresh_Token},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}
