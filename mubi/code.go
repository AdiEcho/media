package mubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (c *LinkCode) New() error {
   req, err := http.NewRequest("", "https://api.mubi.com/v3/link_code", nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   c.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (c *LinkCode) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.LinkCode)
   return b.String()
}

type LinkCode struct {
   AuthToken string `json:"auth_token"`
   LinkCode string `json:"link_code"`
   Raw []byte `json:"-"`
}

func (c *LinkCode) Unmarshal() error {
   return json.Unmarshal(c.Raw, c)
}
