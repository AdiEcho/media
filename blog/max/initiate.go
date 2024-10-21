package max

import (
   "encoding/json"
   "net/http"
)

type link_initiate struct {
   Data struct {
      Attributes struct {
         LinkingCode string
         TargetUrl string
      }
   }
}

func (b *bolt_token) initiate() (*link_initiate, error) {
   req, err := http.NewRequest(
      "POST", prd_api + "/authentication/linkDevice/initiate", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "cookie": {"st=" + b.st},
      "x-device-info": {device_info},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   link := &link_initiate{}
   err = json.NewDecoder(resp.Body).Decode(link)
   if err != nil {
      return nil, err
   }
   return link, nil
}
