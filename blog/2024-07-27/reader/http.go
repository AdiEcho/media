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
   body io.ReadCloser
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   r.body = resp.Body
   return nil
}

func (r response) get_body() io.ReadCloser {
   return r.body
}

func (r *response) set_body(body io.Reader) error {
   return json.NewDecoder(body).Decode(r)
}
