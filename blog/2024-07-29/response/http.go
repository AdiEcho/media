package http

import (
   "encoding/json"
   "io"
   "net/http"
   "time"
)

func pointer[T any](value *T) *T {
   return new(T)
}

type response struct {
   date *time.Time
   body *struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
   raw struct {
      body []byte
      date string
   }
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.raw.date = resp.Header.Get("date")
   r.raw.body, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r *response) unmarshal() error {
   date, err := time.Parse(time.RFC1123, r.raw.date)
   if err != nil {
      return err
   }
   r.date = &date
   r.body = pointer(r.body)
   return json.Unmarshal(r.raw.body, r.body)
}
