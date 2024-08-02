package http

import (
   "encoding/json"
   "net/http"
   "time"
)

func (r *response) unmarshal(text []byte) error {
   return json.Unmarshal(text, r)
}

func (r response) marshal() ([]byte, error) {
   return json.Marshal(r)
}

type response struct {
   Date time.Time
   Body struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
}

func (r *response) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   r.Date, err = time.Parse(time.RFC1123, resp.Header.Get("date"))
   if err != nil {
      return err
   }
   return json.NewDecoder(resp.Body).Decode(&r.Body)
}
