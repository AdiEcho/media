package ctv

import (
   "net/http"
   "strconv"
)

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

// wikipedia.org/wiki/Geo-blocking
func (a axis_content) manifest(m *media_content) string {
   b := []byte("https://capi.9c9media.com/destinations/")
   b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
   b = append(b, "/platforms/desktop/playback/contents/"...)
   b = strconv.AppendInt(b, a.AxisId, 10)
   b = append(b, "/contentPackages/"...)
   b = strconv.AppendInt(b, m.ContentPackages[0].ID, 10)
   b = append(b, "/manifest.mpd"...)
   return string(b)
}
