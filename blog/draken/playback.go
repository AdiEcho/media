package draken

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

func (a auth_login) playback(
   movie *full_movie, title *entitlement,
) (*playback, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.ID
   magine_accesstoken.set(req.Header)
   magine_play_devicemodel.set(req.Header)
   magine_play_deviceplatform.set(req.Header)
   req.Header.Set("authorization", "Bearer " + a.v.Token)
   req.Header.Set("magine-play-deviceid", "!")
   req.Header.Set("magine-play-devicetype", "web")
   req.Header.Set("magine-play-drm", "widevine")
   req.Header.Set("magine-play-entitlementid", title.Token)
   req.Header.Set("magine-play-protocol", "dashs")
   x_forwarded_for.set(req.Header)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   play := new(playback)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type header struct {
   key string
   value string
}

func (h header) set(head http.Header) {
   head.Set(h.key, h.value)
}

var (
   magine_accesstoken = header{
      "magine-accesstoken", "22cc71a2-8b77-4819-95b0-8c90f4cf5663",
   }
   magine_play_devicemodel = header{
      "magine-play-devicemodel", "firefox 111.0 / windows 10",
   }
   magine_play_deviceplatform = header{
      "magine_play_deviceplatform", "firefox",
   }
   x_forwarded_for = header{
      "x-forwarded-for", "78.64.0.0",
   }
)
type playback struct {
   Playlist string
}
