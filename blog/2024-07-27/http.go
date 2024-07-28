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

func (h *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   h.body = resp.Body
   return nil
}

func (h response) write(body io.Writer) (int64, error) {
   defer h.body.Close()
   return io.Copy(body, h.body)
}

func (h *response) read(body io.Reader) error {
   return json.NewDecoder(body).Decode(h)
}

func (h *response) read_body() error {
   defer h.body.Close()
   return json.NewDecoder(h.body).Decode(h)
}
