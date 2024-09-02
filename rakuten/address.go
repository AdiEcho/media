package rakuten

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var classification_id = map[string]int{
   "dk": 283,
   "fi": 284,
   "fr": 23,
   "ie": 41,
   "it": 36,
   "no": 286,
   "pt": 64,
   "se": 282,
   "ua": 276,
   "uk": 18,
}

func (a *Address) String() string {
   var b strings.Builder
   if a.MarketCode != "" {
      b.WriteString("https://www.rakuten.tv/")
      b.WriteString(a.MarketCode)
   }
   if a.ContentId != "" {
      b.WriteString("/movies/")
      b.WriteString(a.ContentId)
   }
   return b.String()
}

func (a *Address) video(quality string) *OnDemand {
   var o OnDemand
   o.AudioLanguage = "ENG"
   o.AudioQuality = "2.0"
   o.ContentType = "movies"
   o.DeviceSerial = "!"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
   o.DeviceIdentifier = "atvui40"
   o.ClassificationId = a.ClassificationId
   o.ContentId = a.ContentId
   o.DeviceStreamVideoQuality = quality
   return &o
}

func (a *Address) Hd() *OnDemand {
   return a.video("HD")
}

func (a *Address) Fhd() *OnDemand {
   return a.video("FHD")
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "rakuten.tv")
   s = strings.TrimPrefix(s, "/")
   var found bool
   a.MarketCode, a.ContentId, found = strings.Cut(s, "/movies/")
   if !found {
      return errors.New("/movies/ not found")
   }
   a.ClassificationId, found = classification_id[a.MarketCode]
   if !found {
      return errors.New("MarketCode not found")
   }
   return nil
}

type Address struct {
   ClassificationId int
   ContentId        string
   MarketCode       string
}

func (a *Address) Movie() (*GizmoMovie, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.ContentId
   req.URL.RawQuery = url.Values{
      "market_code":       {a.MarketCode},
      "classification_id": {strconv.Itoa(a.ClassificationId)},
      "device_identifier": {"atvui40"},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   movie := &GizmoMovie{}
   err = json.NewDecoder(resp.Body).Decode(movie)
   if err != nil {
      return nil, err
   }
   return movie, nil
}
