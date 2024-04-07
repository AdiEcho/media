package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type anonymous struct {
   AuthToken string
}

func (a *anonymous) New() error {
   req, err := http.NewRequest(
      "POST", "https://plex.tv/api/v2/users/anonymous", nil,
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "X-Plex-Product": {"Plex Mediaverse"},
      "X-Plex-Client-Identifier": {"!"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a anonymous) abs(path string, query url.Values) string {
   query.Set("x-plex-token", a.AuthToken)
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = path
   u.RawQuery = query.Encode()
   u.Scheme = "https"
   return u.String()
}
