package sbs

import (
   "net/http"
   "net/url"
)

func auth_native(user, pass string) (*http.Response, error) {
   return http.PostForm(
      "https://www.sbs.com.au/api/v3/janrain/auth_native_traditional",
      url.Values{
         "pass": {pass},
         "user": {user},
      },
   )
}
