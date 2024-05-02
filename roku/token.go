package roku

import (
   "encoding/json"
   "net/http"
)

const user_agent = "trc-googletv; production; 0"

type account_token struct {
   AuthToken string
}

func (a *account_token) New() error {
   req, err := http.NewRequest(
      "GET", "https://googletv.web.roku.com/api/v1/account/token", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("user-agent", user_agent)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
