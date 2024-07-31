package http

import (
   "encoding/json"
   "io"
   "net/http"
)

func pointer[T any](value *T) *T {
   return new(T)
}

type response struct {
   fly_request_id string
   body *struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
   raw_body []byte
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.fly_request_id = resp.Header.Get("fly-request-id")
   r.raw_body, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r *response) unmarshal_body() error {
   r.body = pointer(r.body)
   return json.Unmarshal(r.raw_body, r.body)
}
