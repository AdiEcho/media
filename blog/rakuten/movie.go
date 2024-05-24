package rakuten

import (
   "net/http"
   "net/url"
)

func gizmo_movie() (*http.Response, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/jerry-maguire"
   req.URL.RawQuery = url.Values{
      "classification_id": {"23"},
      "market_code": {"fr"},
   }.Encode()
   return http.DefaultClient.Do(req)
}
