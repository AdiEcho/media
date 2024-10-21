package max

import "net/http"

const (
   device_info = "beam/1.1.2.2 (LG/OLED55C9PVA; webOS/4.9.0-05.00.03; 04137aa2-1e1e-6f52-7a08-12249c864690/9e1b83a7-ddde-42c9-b335-a54232bd2a9f)"
   prd_api = "https://default.prd.api.discomax.com"
)

type bolt_token struct {
   session_state string
   st *http.Cookie
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
   b.session_state = resp.Header.Get("x-wbd-session-state")
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         b.st = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}
