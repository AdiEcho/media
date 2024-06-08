package rtbf

import (
   "encoding/json"
   "net/http"
)

func (auvio_page) Show() string {
   return ""
}

func (auvio_page) Season() int {
   return 0
}

func (auvio_page) Episode() int {
   return 0
}

func (auvio_page) Title() string {
   return ""
}

func (auvio_page) Year() int {
   return 0
}

type auvio_page struct {
   Data struct {
      Content struct {
         Program *struct {
            Title string
         }
         Subtitle string
         Title string
      }
   }
}

func (a *auvio_page) New(path string) error {
   res, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
