package http

import (
   "encoding/json"
   "io"
   "net/http"
)

func pointer[T any](value *T) *T {
   return new(T)
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

type response struct {
   body *struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
   raw []byte
}

func (r response) unmarshal() error {
   r.body = pointer(r.body)
   return json.Unmarshal(r.raw, r.body)
}
