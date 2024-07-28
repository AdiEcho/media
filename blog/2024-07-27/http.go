package http

import (
   "encoding/json"
   "io"
   "net/http"
)

type response struct {
   Slideshow *struct {
      Date string
      Title string
   }
   raw []byte
}

func (h *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   h.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (h response) marshal() []byte {
   return h.raw
}

func (h *response) unmarshal(raw []byte) error {
   return json.Unmarshal(raw, h)
}

func (h *response) unmarshal_raw() error {
   return h.unmarshal(h.raw)
}
