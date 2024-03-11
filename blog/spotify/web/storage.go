package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type storage_resolve struct {
   CDNURL []string
}

func (s *storage_resolve) New(login android.LoginOk, file_id string) error {
   token, err := login.AccessToken()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("GET", "https://guc3-spclient.spotify.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/storage-resolve/v2/files/audio/interactive/10/" + file_id
   req.Header.Set("authorization", "Bearer " + token)
   req.URL.RawQuery = "alt=json"
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
   return json.NewDecoder(res.Body).Decode(s)
}
