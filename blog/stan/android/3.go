package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type app_session struct {
   JwToken string
}

func (w web_token) session() (*app_session, error) {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/sessions/mobile/app", url.Values{
         "jwToken": {w.v.JwToken},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   session := new(app_session)
   if err := json.NewDecoder(res.Body).Decode(session); err != nil {
      return nil, err
   }
   return session, nil
}
