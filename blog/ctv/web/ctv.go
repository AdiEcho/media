package ctv

import (
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

type content_packages struct {
   data []byte
   v struct {
      Items []struct {
         ID int64
      }
   }
}

func (c *content_packages) unmarshal() error {
   return json.Unmarshal(c.data, &c.v)
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

type last_segment struct {
   Content struct {
      AxisId int64
      FirstPlayableContent struct {
         AxisId int64
      }
   }
}

// wikipedia.org/wiki/Geo-blocking
func (s last_segment) manifest(c content_packages) string {
   b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
   b = append(b, "/platforms/desktop/playback/contents/"...)
   b = strconv.AppendInt(b, s.Content.FirstPlayableContent.AxisId, 10)
   b = append(b, "/contentPackages/"...)
   b = strconv.AppendInt(b, c.v.Items[0].ID, 10)
   b = append(b, "/manifest.mpd"...)
   return string(b)
}

func (s last_segment) packages() (*content_packages, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, s.Content.FirstPlayableContent.AxisId, 10)
      b = append(b, "/contentPackages"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var packages content_packages
   packages.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &packages, nil
}
