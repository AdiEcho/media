package mubi

import (
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

func (c linkCode) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.s.Link_Code)
   return b.String()
}

type linkCode struct {
   Raw []byte
   s struct {
      Auth_Token string
      Link_Code string
   }
}

func (c *linkCode) unmarshal() error {
   return json.Unmarshal(c.Raw, &c.s)
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
   c.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}
