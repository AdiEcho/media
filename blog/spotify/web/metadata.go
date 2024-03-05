package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type metadata struct {
   File []struct {
      File_ID string
      Format string
   }
}

func (m *metadata) New(login android.LoginOk, track string) error {
   token, ok := login.AccessToken()
   if !ok {
      return errors.New("android.LoginOk.AccessToken")
   }
   req, err := http.NewRequest("GET", "https://spclient.wg.spotify.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/metadata/4/track/" + track
   req.URL.RawQuery = "market=from_token"
   req.Header = http.Header{
      "Accept": {"application/json"},
      "Authorization": {"Bearer " + token},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(res.Body).Decode(m)
}
