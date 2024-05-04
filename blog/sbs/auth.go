package sbs

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type auth_native struct {
   User_State struct {
      SessionToken string
   }
}

func (a *auth_native) New(user, pass string) error {
   res, err := http.PostForm(
      "https://www.sbs.com.au/api/v3/janrain/auth_native_traditional",
      url.Values{
         "pass": {pass},
         "user": {user},
         "express": {"1"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
