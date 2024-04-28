package ctv

import (
   "net/http"
   "strconv"
)

// wikipedia.org/wiki/Geo-blocking
func manifest(axis_id, content_package int64) string {
   b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
   b = append(b, "/platforms/desktop/playback/contents/"...)
   b = strconv.AppendInt(b, axis_id, 10)
   b = append(b, "/contentPackages/"...)
   b = strconv.AppendInt(b, content_package, 10)
   b = append(b, "/manifest.mpd"...)
   return string(b)
}

type poster struct{}

func (poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (poster) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) RequestUrl() (string, bool) {
   return "https://license.9c9media.ca/widevine", true
}
