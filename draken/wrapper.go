package draken

import (
   "bytes"
   "io"
   "net/http"
)

type Wrapper func() (*AuthLogin, *Playback)

func (w Wrapper) Wrap(data []byte) ([]byte, error) {
   auth, play := w()
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/playback/v1/widevine/license",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   magine_accesstoken.set(req.Header)
   for key, value := range play.Headers {
      req.Header.Set(key, value)
   }
   req.Header.Set("authorization", "Bearer " + auth.Token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
