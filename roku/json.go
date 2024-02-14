package roku

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func (c *CrossSite) New() error {
   // this has smaller body than www.roku.com
   res, err := http.Get("https://therokuchannel.roku.com")
   if err != nil {
      return err
   }
   defer res.Body.Close()
   c.cookies = res.Cookies()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   _, text = json.Cut(text, []byte("\tcsrf:"), nil)
   return json.Decode(text, &c.token)
}
