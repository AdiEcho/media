package http

import (
   "encoding/json"
   "io"
   "net/http"
)

type response struct {
   header struct {
      fly_request_id string
   }
   body struct {
      Slideshow *struct {
         Date string
         Title string
      }
      raw []byte
   }
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.header.fly_request_id = resp.Header.Get("fly-request-id")
   r.body.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r response) get_body() []byte {
   return r.body.raw
}

func (r *response) set_body(raw []byte) error {
   return json.Unmarshal(raw, &r.body)
}
