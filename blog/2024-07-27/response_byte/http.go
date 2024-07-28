package http

import (
   "encoding/json"
   "io"
   "net/http"
)

type response struct {
   Slideshow struct {
      Date string
      Title string
   }
   response *http.Response
}

func (h *response) New() error {
   var err error
   h.response, err = http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   return nil
}

func (h response) marshal() ([]byte, error) {
   defer h.response.Body.Close()
   return io.ReadAll(h.response.Body)
}

func (h *response) unmarshal(text []byte) error {
   return json.Unmarshal(text, h)
}
