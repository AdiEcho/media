package stan

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
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
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   session := new(app_session)
   if err := json.NewDecoder(res.Body).Decode(session); err != nil {
      return nil, err
   }
   return session, nil
}
