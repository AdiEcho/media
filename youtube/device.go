package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func New_Device_Code() (*Device_Code, error) {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/device/code",
      url.Values{
         "client_id": {client_ID},
         "scope": {"https://www.googleapis.com/auth/youtube"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   code := new(Device_Code)
   if err := json.NewDecoder(res.Body).Decode(code); err != nil {
      return nil, err
   }
   return code, nil
}
