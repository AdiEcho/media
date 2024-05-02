package roku

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.DRM.Widevine.LicenseServer, true
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (a account_token) playback(roku_id string) (*playback, error) {
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(playback)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}
const user_agent = "trc-googletv; production; 0"

type account_token struct {
   AuthToken string
}

func (a *account_token) New() error {
   req, err := http.NewRequest(
      "GET", "https://googletv.web.roku.com/api/v1/account/token", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("user-agent", user_agent)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
