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
   body []byte
   fly_request_id string
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.fly_request_id = resp.Header.Get("fly-request-id")
   r.body, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r response) get_body() []byte {
   return r.body
}

func (r *response) set_body(body []byte) error {
   return json.Unmarshal(body, r)
}
