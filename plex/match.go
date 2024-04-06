package plex

import (
   "net/http"
   "net/url"
)

func (a anonymous) matches(path string) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {path},
      "x-plex-token": {a.AuthToken},
   }.Encode()
   return http.DefaultClient.Do(req)
}
