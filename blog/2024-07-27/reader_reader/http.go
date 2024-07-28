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

func (h response) write(dst io.Writer) (int64, error) {
   defer h.body.Close()
   return io.Copy(dst, h.body)
}

func (h *response) read(src io.Reader) error {
   return json.NewDecoder(src).Decode(h)
}
