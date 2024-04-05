package plex

import "net/http"

func anonymous() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "https://plex.tv/api/v2/users/anonymous", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "X-Plex-Product": {"Plex Mediaverse"},
      "X-Plex-Client-Identifier": {"!"},
   }
   return http.DefaultClient.Do(req)
}
