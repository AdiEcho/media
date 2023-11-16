package hulu

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type Authenticate struct {
   Data struct {
      User_Token string
   }
}

func Living_Room(user, password string) (*Authenticate, error) {
   res, err := http.PostForm(
      "https://auth.hulu.com/v2/livingroom/password/authenticate", url.Values{
         "friendly_name": {"!"},
         "password": {password},
         "serial_number": {"!"},
         "user_email": {user},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Authenticate)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}
