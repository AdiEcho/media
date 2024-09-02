package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type AccountAuth struct {
   AuthToken string
   Raw []byte `json:"-"`
}

// token can be nil
func (a *AccountAuth) New(token *AccountToken) error {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/api/v1/account/token"
   req.Header.Set("user-agent", user_agent)
   if token != nil {
      req.Header.Set("x-roku-content-token", token.V.Token)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   a.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AccountAuth) Unmarshal() error {
   a.v = pointer(a.v)
   return json.Unmarshal(a.Data, a.v)
}

func (a AccountAuth) Playback(roku_id string) (*Playback, error) {
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
      "x-roku-content-token": {a.v.AuthToken},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   play := &Playback{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}
