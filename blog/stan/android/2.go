package stan

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type web_token struct {
   JwToken string
   ProfileId string
}

func (a activation_code) token() (*web_token, error) {
   res, err := http.Get(a.v.URL)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   web := new(web_token)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}
