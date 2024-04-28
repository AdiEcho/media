package ctv

import (
   "net/http"
   "strconv"
)

func new_content_packages(axis_id int64) (*http.Response, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, axis_id, 10)
      b = append(b, "/contentPackages"...)
      return string(b)
   }()
   return http.Get(address)
}
