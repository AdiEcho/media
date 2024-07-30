package http

import (
   "encoding/json"
   "io"
   "net/http"
)

type response_body struct {
   Slideshow *struct {
      Date string
      Title string
   }
   raw []byte
}

func (r *response_body) New() error {
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

func (r response_body) get() []byte {
   return r.raw
}

func (r *response_body) set(raw []byte) error {
   return json.Unmarshal(raw, r)
}
