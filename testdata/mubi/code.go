package mubi

import (
   "encoding/json"
   "net/http"
)

type linkCode struct {
   Auth_Token string
   Link_Code string
}

func (c *linkCode) New(country string) error {
   req, err := http.NewRequest("GET", "https://api.mubi.com/v3/link_code", nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Client": {"android"},
      "Client-Country": {country},
      "Client-Device-Identifier": {"!"},
      "Client-Version": {"!"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(c)
}
