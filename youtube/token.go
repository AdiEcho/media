package youtube

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
)

// YouTube on TV
const (
   client_id = "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func (d DeviceCode) Token() (*OauthToken, error) {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token", url.Values{
         "client_id": {client_id},
         "client_secret": {client_secret},
         "device_code": {d.Device_Code},
         "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var token OauthToken
   token.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &token, nil
}

type OauthToken struct {
   Raw []byte
   Token struct {
      Access_Token string
      Refresh_Token string
   }
}

func (o *OauthToken) Unmarshal() error {
   return json.Unmarshal(o.Raw, &o.Token)
}

func (o *OauthToken) Refresh() error {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_id},
         "client_secret": {client_secret},
         "grant_type": {"refresh_token"},
         "refresh_token": {o.Token.Refresh_Token},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(&o.Token)
}
