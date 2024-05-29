package roku

import (
   "encoding/json"
   "net/http"
)

const user_agent = "Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"

type account_token struct {
   AuthToken string
}

// token can be nil
func (a *account_token) New(token *activation_token) error {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/api/v1/account/token"
   req.Header.Set("user-agent", user_agent)
   if token != nil {
      req.Header.Set("x-roku-content-token", token.v.Token)
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
