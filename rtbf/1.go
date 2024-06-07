package rtbf

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type embed_media struct {
   Data struct {
      AssetId string
      Program struct {
         Title string
      }
      Title string
   }
   Meta struct {
      SmartAds struct {
         // these can be empty string, so we cannot use these:
         // int `json:",string"`
         // json.Number
         CTE string
         CTS string
      }
   }
}

func (embed_media) Show() string {
   return ""
}

func (embed_media) Season() int {
   return 0
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

func (embed_media) Episode() int {
   return 0
}

func (e embed_media) Title() string {
   return e.Data.Title
}

// its just not available from what I can tell
func (embed_media) Year() int {
   return 0
}
