package draken

import (
   "net/http"
   "net/url"
   "os"
)

func (a auth_login) playback(
   title entitlement, movie full_movie,
) (*http.Response, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.ID
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Token},
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
   return http.DefaultClient.Do(req)
}
