package hulu

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
)

type Authenticate struct {
   Raw []byte
   Value struct {
      Data struct {
         User_Token string
      }
   }
}

func Living_Room(email, password string) (*Authenticate, error) {
   res, err := http.PostForm(
      "https://auth.hulu.com/v2/livingroom/password/authenticate", url.Values{
         "friendly_name": {"!"},
         "password": {password},
         "serial_number": {"!"},
         "user_email": {email},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var auth Authenticate
   auth.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Raw, &a.Value)
}
