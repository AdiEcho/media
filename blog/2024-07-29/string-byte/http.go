package http

import (
   "encoding/json"
   "io"
   "net/http"
   "time"
)

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.date.raw = resp.Header.Get("date")
   r.body.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func pointer[T any](value *T) *T {
   return new(T)
}

type response struct {
   date struct {
      value *time.Time
      raw string
   }
   body struct {
      value *struct {
         Slideshow struct {
            Date string
            Title string
         }
      }
      raw []byte
   }
}

func (r *response) unmarshal() error {
   date, err := time.Parse(time.RFC1123, r.date.raw)
   if err != nil {
      return err
   }
   r.date.value = &date
   r.body.value = pointer(r.body.value)
   return json.Unmarshal(r.body.raw, r.body.value)
}
