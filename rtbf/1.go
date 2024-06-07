package rtbf

import (
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

func (n *number) UnmarshalText(text []byte) error {
   if len(text) >= 1 {
      i, err := strconv.Atoi(string(text))
      if err != nil {
         return err
      }
      *n = number(i)
   }
   return nil
}

type embed_media struct {
   Data struct {
      AssetId string
      Program *struct {
         Title string
      }
      Subtitle string
      Title string
   }
   Meta struct {
      SmartAds struct {
         CTE number
         CTS number
      }
   }
}

func (e embed_media) Episode() int {
   return int(e.Meta.SmartAds.CTE)
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

func (e embed_media) Season() int {
   return int(e.Meta.SmartAds.CTS)
}

func (e embed_media) Show() string {
   if v := e.Data.Program; v != nil {
      return v.Title
   }
   return ""
}

func (e embed_media) Title() string {
   if e.Data.Program != nil {
      // json.data.subtitle = "06 - Les ombres de la guerre";
      _, after, _ := strings.Cut(e.Data.Subtitle, " - ")
      return after
   }
   // json.data.title = "I care a lot";
   return e.Data.Title
}

// its just not available from what I can tell
func (embed_media) Year() int {
   return 0
}

type number int
