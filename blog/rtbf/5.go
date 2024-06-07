package rtbf

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (w web_token) four() (*http.Response, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Device struct {
            DeviceId string `json:"deviceId"`
            Name string `json:"name"`
            Type string `json:"type"`
         } `json:"device"`
         JWT string `json:"jwt"`
      }
      s.Device.DeviceId = "7f5cdd55-1cfe-4841-9e8e-ecd8b823cfad"
      s.Device.Name = "Browser"
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
   req.Header = http.Header{
      "accept": {"application/json"},
      "accept-language": {"en-US,en;q=0.5"},
      "content-type": {"application/json"},
      "origin": {"https://auvio.rtbf.be"},
      "referer": {"https://auvio.rtbf.be/"},
      "sec-fetch-dest": {"empty"},
      "sec-fetch-mode": {"cors"},
      "sec-fetch-site": {"cross-site"},
      "te": {"trailers"},
      "user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"},
   }
   return http.DefaultClient.Do(req)
}
