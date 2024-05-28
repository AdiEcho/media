package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

type AccountToken struct {
   AuthToken string
}

func (a *AccountToken) New() error {
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

func (a AccountToken) Playback(roku_id string) (*Playback, error) {
   body, err := json.Marshal(map[string]string{
      "mediaFormat": "DASH",
      "providerId":  "rokuavod",
      "rokuId":      roku_id,
   })
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
      "content-type":         {"application/json"},
      "user-agent":           {user_agent},
      "x-roku-content-token": {a.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Playback)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
   URL string
}

const user_agent = "trc-googletv; production; 0"

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
