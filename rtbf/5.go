package rtbf

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type gigya_login struct {
   SessionToken string
}

func (w web_token) login() (*gigya_login, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Device struct {
            DeviceId string `json:"deviceId"`
            Type string `json:"type"`
         } `json:"device"`
         JWT string `json:"jwt"`
      }
      s.Device.Type = "WEB"
      s.JWT = w.IdToken
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://exposure.api.redbee.live", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin"
   req.Header.Set("content-type", "application/json")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(gigya_login)
   err = json.NewDecoder(res.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}
