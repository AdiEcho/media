package rtbf

import (
   "net/http"
   "strconv"
)

func one(media int64) (*http.Response, error) {
   address := func() string {
      b := []byte("https://bff-service.rtbf.be/auvio/v1.23/embed/media/")
      b = strconv.AppendInt(b, media, 10)
      b = append(b, "?userAgent"...)
      return string(b)
   }()
   return http.Get(address)
}
