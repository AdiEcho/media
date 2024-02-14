package roku

import (
   "154.pages.dev/encoding/json"
   "errors"
   "io"
   "net/http"
)

func NewCrossSite() (*CrossSite, error) {
   // this has smaller body than www.roku.com
   res, err := http.Get("https://therokuchannel.roku.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var site CrossSite
   site.cookies = res.Cookies()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = json.Cut(text, []byte("\tcsrf:"), nil)
   if err := json.Decode(text, &site.token); err != nil {
      return nil, err
   }
   return &site, nil
}
