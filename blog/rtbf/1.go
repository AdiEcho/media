package rtbf

import "net/http"

type one struct {
   cookie *http.Cookie
}

func (o *one) New() error {
   res, err := http.Get("https://login.auvio.rtbf.be/accounts.webSdkBootstrap")
   if err != nil {
      return err
   }
   defer res.Body.Close()
   for _, cookie := range res.Cookies() {
      if cookie.Name == "gmid" {
         o.cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}
