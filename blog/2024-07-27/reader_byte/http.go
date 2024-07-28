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
   body io.ReadCloser
}

func (h *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   h.body = resp.Body
   return nil
}

func (h response) marshal() ([]byte, error) {
   defer h.body.Close()
   return io.ReadAll(h.body)
}

func (h *response) unmarshal(text []byte) error {
   return json.Unmarshal(text, h)
}
