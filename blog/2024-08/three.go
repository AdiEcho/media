package http

import (
   "encoding/json"
   "net/http"
   "time"
)

func (t *three) unmarshal(text []byte) error {
   return json.Unmarshal(text, t)
}

func (t three) marshal() ([]byte, error) {
   return json.Marshal(t)
}

type three struct {
   Date time.Time
   Body struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
}

func (t *three) New() error {
   resp, err := http.Get("http://httpbingo.org/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   t.Date, err = time.Parse(time.RFC1123, resp.Header.Get("date"))
   if err != nil {
      return err
   }
   return json.NewDecoder(resp.Body).Decode(&t.Body)
}
