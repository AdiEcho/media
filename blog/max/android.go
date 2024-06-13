package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

type default_token struct {
   Data struct {
      Attributes struct {
         Token string
      }
   }
}

func (d *default_token) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.max.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set(
      "x-device-info",
      "BEAM-Android/4.1.1 (unknown/Android SDK built for x86; ANDROID/6.0; bafeab496eee63aa/b6746ddc-7bc7-471f-a16c-f6aaf0c34d26)",
   )
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(resp.Body).Decode(d)
}
