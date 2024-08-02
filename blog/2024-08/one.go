package http

import (
   "encoding/json"
   "io"
   "net/http"
   "time"
)

type one struct {
   date value[time.Time]
   body value[struct {
      Slideshow struct {
         Date string
         Title string
      }
   }]
}

func (o *one) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   o.date.raw = []byte(resp.Header.Get("date"))
   o.body.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (o *one) unmarshal() error {
   date, err := time.Parse(time.RFC1123, string(o.date.raw))
   if err != nil {
      return err
   }
   o.date.value = &date
   o.body.New()
   return json.Unmarshal(o.body.raw, o.body.value)
}

type value[T any] struct {
   value *T
   raw []byte
}

func (v *value[T]) New() {
   v.value = new(T)
}
