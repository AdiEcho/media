package roku

import (
   "encoding/json"
   "io"
   "net/http"
)

type activation_token struct {
   data []byte
   v struct {
      Token string
   }
}

func (a activation_code) token() (*activation_token, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + a.V.Code
   req.Header = http.Header{
      "user-agent": {user_agent},
      "x-roku-content-token": {a.Token.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var token activation_token
   token.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &token, nil
}

func (a *activation_token) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
