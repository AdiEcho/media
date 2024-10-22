package max

import (
   "encoding/json"
   "io"
   "net/http"
)

type LinkLogin struct {
   State string
   Raw []byte
   token string
}

// you must
// /authentication/linkDevice/initiate
// first or this will always fail
func (b BoltToken) Login() (*LinkLogin, error) {
   req, err := http.NewRequest(
      "POST", prd_api + "/authentication/linkDevice/login", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("cookie", "st=" + b.st)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var link LinkLogin
   link.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   link.state = resp.Header.Get("x-wbd-session-state")
   return &link, nil
}

func (v *LinkLogin) Unmarshal() error {
   var value struct {
      Data struct {
         Attributes struct {
            Token string
         }
      }
   }
   err := json.Unmarshal(v.raw, &value)
   if err != nil {
      return err
   }
   v.token = value.Data.Attributes.Token
   return nil
}
