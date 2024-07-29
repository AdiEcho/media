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

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r response) marshal() []byte {
   return r.raw
}

func (r *response) unmarshal(raw []byte) error {
   return json.Unmarshal(raw, r)
}
