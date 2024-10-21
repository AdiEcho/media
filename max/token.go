package max

import "net/http"

const (
   device_info = "!/!(!/!;!/!;!/!)"
   prd_api = "https://default.prd.api.discomax.com"
)

type bolt_token struct {
   st string
}

func (b *bolt_token) New() error {
   req, err := http.NewRequest("", prd_api + "/token?realm=bolt", nil)
   if err != nil {
      return err
   }
   req.Header.Set("x-device-info", device_info)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         b.st = cookie.Value
         return nil
      }
   }
   return http.ErrNoCookie
}
