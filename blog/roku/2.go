package roku

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (a account_token) playback(roku_id string) (*http.Response, error) {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "mediaFormat": "DASH",
         "rokuId": roku_id,
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v3/playback",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {user_agent},
      "X-Roku-Content-Token": {a.AuthToken},
   }
   return http.DefaultClient.Do(req)
}
