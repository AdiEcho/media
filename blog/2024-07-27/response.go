package hello

import (
   "encoding/json"
   "io"
   "net/http"
)

func (h http_bin) marshal() ([]byte, error) {
   defer h.response.Body.Close()
   return io.ReadAll(h.response.Body)
}

func (h *http_bin) unmarshal(text []byte) error {
   return json.Unmarshal(text, h)
}

type http_bin struct {
   Slideshow struct {
      Date string
      Title string
   }
   response *http.Response
}

func (h *http_bin) New() error {
   var err error
   h.response, err = http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   return nil
}
