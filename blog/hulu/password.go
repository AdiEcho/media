package hulu

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type authenticate struct {
   Data struct {
      User_Token string
   }
}

// github.com/matthuisman/slyguy.addons/blob/master/slyguy.hulu/resources/lib/api.py
func living_room(user, password string) (*authenticate, error) {
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
   auth := new(authenticate)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}
