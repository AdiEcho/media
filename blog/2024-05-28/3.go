package roku

import (
   "encoding/json"
   "net/http"
)

type three_response struct {
   Token string
}

func (t two_response) three() (*three_response, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + t.Two.Code
   req.Header = http.Header{
      "user-agent": {user_agent},
      "x-roku-content-token": {t.One.AuthToken},
   }
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["X-Requested-With"] = []string{"com.roku.web.trc"}
   req.Header["X-Roku-Code-Version"] = []string{"2"}
   req.Header["X-Roku-Reserved-Channel-Store-Code"] = []string{"us"}
   req.Header["X-Roku-Reserved-Culture-Code"] = []string{"en-US"}
   req.Header["X-Roku-Reserved-Experiment-Configs"] = []string{"e30="}
   req.Header["X-Roku-Reserved-Experiment-State"] = []string{"W10="}
   req.Header["X-Roku-Reserved-Session-Id"] = []string{"f77813e1-a689-41e9-b058-097e8520f4d2"}
   req.Header["X-Roku-Reserved-Time-Zone-Offset"] = []string{"+00:00"}
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   three := new(three_response)
   err = json.NewDecoder(res.Body).Decode(three)
   if err != nil {
      return nil, err
   }
   return three, nil
}
