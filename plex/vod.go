package plex

import "net/http"

func (a anonymous) vod(m *metadata) (*http.Response, error) {
   req, err := http.NewRequest("GET", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/library/metadata/" + m.RatingKey
   req.Header = http.Header{
      "accept": {"application/json"},
      "x-plex-token": {a.AuthToken},
   }
   return http.DefaultClient.Do(req)
}
