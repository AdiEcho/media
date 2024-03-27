package stan

import (
   "net/http"
   "net/url"
)

func (w web_token) session() (*http.Response, error) {
   return http.PostForm(
      "https://api.stan.com.au/login/v1/sessions/mobile/app", url.Values{
         "jwToken":[]string{w.JwToken},
      },
   )
}
