package rtbf

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type embed_media struct {
   Data struct {
      AssetId string
   }
}

func (e *embed_media) New(media int64) error {
   address := func() string {
      b := []byte("https://bff-service.rtbf.be/auvio/v1.23/embed/media/")
      b = strconv.AppendInt(b, media, 10)
      b = append(b, "?userAgent"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(e)
}
