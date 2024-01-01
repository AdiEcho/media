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

func (t *Token) Refresh() error {
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

func (d Device_Code) Token() (*Token, error) {
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
      return nil, err
   }
   defer res.Body.Close()
   tok := new(Token)
   if err := json.NewDecoder(res.Body).Decode(tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func Read_Token(name string) (*Token, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   tok := new(Token)
   if err := json.Unmarshal(text, tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func (t Token) Write_File(name string) error {
   text, err := json.Marshal(t)
   if err != nil {
      return err
   }
   return os.WriteFile(name, text, 0666)
}

type Token struct {
   Access_Token string
   Error string
   Refresh_Token string
}
