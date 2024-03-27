package stan

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type web_token struct {
   data []byte
   v struct {
      JwToken string
      ProfileId string
   }
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
   var web web_token
   web.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &web, nil
}

func (w *web_token) unmarshal() error {
   return json.Unmarshal(w.data, &w.v)
}
