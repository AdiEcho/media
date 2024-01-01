package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (d *Device_Code) Post() error {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/device/code",
      url.Values{
         "client_id": {client_ID},
         "scope": {"https://www.googleapis.com/auth/youtube"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(d)
}
