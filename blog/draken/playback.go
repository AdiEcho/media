package draken

import (
   "encoding/json"
   "net/http"
)

type playback struct {
   Playlist string
}

func (a auth_login) playback(
   movie *full_movie, title *entitlement,
) (*playback, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.ID
   req.Header = http.Header{
      "authorization": {"Bearer " + a.v.Token},
      "magine-accesstoken": {magine_accesstoken},
      "magine-play-deviceid": {"!"},
      "magine-play-devicemodel": {"firefox 111.0 / windows 10"},
      "magine-play-deviceplatform": {"firefox"},
      "magine-play-devicetype": {"web"},
      "magine-play-drm": {"widevine"},
      "magine-play-entitlementid": {title.Token},
      "magine-play-protocol": {"dashs"},
      "x-forwarded-for": {"78.64.0.0"},
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
