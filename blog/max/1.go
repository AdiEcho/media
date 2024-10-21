package max

import "net/http"

const device_info = "beam/1.1.2.2 (LG/OLED55C9PVA; webOS/4.9.0-05.00.03; 04137aa2-1e1e-6f52-7a08-12249c864690/9e1b83a7-ddde-42c9-b335-a54232bd2a9f)"

func head() (string, error) {
   req, err := http.NewRequest(
      "", "https://default.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return "", err
   }
   req.Header.Set("x-device-info", device_info)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   for _, cookie := resp.Cookies() {
      if cookie.Name == "st" {
         return cookie.Value, nil
      }
   }
   return "", http.ErrNoCookie
}

func body() (string, error) {
   req, err := http.NewRequest(
      "", "https://default.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return "", err
   }
   req.Header.Set("x-device-info", device_info)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   var value struct {
      Attributes struct {
         Token string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return "", err
   }
   return value.Attributes.Token, nil
}
