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
   response *http.Response
}

func (h *response) New() error {
   var err error
   h.response, err = http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   return nil
}

func (h response) write(dst io.Writer) (int64, error) {
   defer h.response.Body.Close()
   return io.Copy(dst, h.response.Body)
}

func (h *response) read(src io.Reader) error {
   return json.NewDecoder(src).Decode(h)
}
