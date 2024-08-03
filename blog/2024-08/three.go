package http

import (
   "encoding/json"
   "net/http"
   "time"
)

type response_three struct {
   Date time.Time
   Body struct {
      Slideshow struct {
         Author string `json:"author"`
         Date   string `json:"date"`
         Slides []struct {
            Title string   `json:"title"`
            Type  string   `json:"type"`
            Items []string `json:"items,omitempty"`
         } `json:"slides"`
         Title string `json:"title"`
      } `json:"slideshow"`
   }
}

func (r *response_three) New() error {
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

func (r response_three) marshal() ([]byte, error) {
   return json.Marshal(r)
}

func (r *response_three) unmarshal(text []byte) error {
   return json.Unmarshal(text, r)
}
