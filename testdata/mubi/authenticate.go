package mubi

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

type authenticate struct {
   s struct {
      Token string
   }
   Raw []byte
}

func (c linkCode) authenticate() (*authenticate, error) {
   body, err := json.Marshal(map[string]string{"auth_token": c.s.Auth_Token})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.mubi.com/v3/authenticate", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Client": {"android"},
      "Client-Country": {client_country},
      "Client-Device-Identifier": {"!"},
      "Client-Version": {"!"},
      "Content-Type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var auth authenticate
   auth.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a *authenticate) unmarshal() error {
   return json.Unmarshal(a.Raw, &a.s)
}
