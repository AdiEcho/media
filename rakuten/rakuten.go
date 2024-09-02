package rakuten

import (
   "bytes"
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

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "rakuten.tv")
   s = strings.TrimPrefix(s, "/")
   var found bool
   a.market_code, a.content_id, found = strings.Cut(s, "/movies/")
   if !found {
      return errors.New("/movies/ not found")
   }
   a.classification_id, found = classification_id[a.market_code]
   if !found {
      return errors.New("market_code not found")
   }
   return nil
}

type Address struct {
   classification_id int
   content_id        string
   market_code       string
}

func (a Address) Movie() (*GizmoMovie, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.content_id
   req.URL.RawQuery = url.Values{
      "market_code":       {a.market_code},
      "classification_id": {strconv.Itoa(a.classification_id)},
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

func (a Address) String() string {
   var b strings.Builder
   if a.market_code != "" {
      b.WriteString("https://www.rakuten.tv/")
      b.WriteString(a.market_code)
   }
   if a.content_id != "" {
      b.WriteString("/movies/")
      b.WriteString(a.content_id)
   }
   return b.String()
}

func (GizmoMovie) Show() string {
   return ""
}

func (GizmoMovie) Season() int {
   return 0
}

func (GizmoMovie) Episode() int {
   return 0
}

func (g GizmoMovie) Title() string {
   return g.Data.Title
}

type GizmoMovie struct {
   Data struct {
      Title string
      Year  int
   }
}

func (g GizmoMovie) Year() int {
   return g.Data.Year
}
// github.com/mitmproxy/mitmproxy/blob/main/mitmproxy/contentviews/protobuf.py
func (StreamInfo) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("content-type", "application/x-protobuf")
   return head, nil
}

func (StreamInfo) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (a Address) Hd() OnDemand {
   return a.video("HD")
}

type StreamInfo struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

// geo block
func (o OnDemand) Info() (*StreamInfo, error) {
   body, err := json.Marshal(o)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gizmo.rakuten.tv/v3/avod/streamings",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var data struct {
      Data struct {
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &data.Data.StreamInfos[0], nil
}

func (a Address) Fhd() OnDemand {
   return a.video("FHD")
}

func (s StreamInfo) RequestUrl() (string, bool) {
   return s.LicenseUrl, true
}

func (StreamInfo) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type OnDemand struct {
   AudioLanguage            string `json:"audio_language"`
   AudioQuality             string `json:"audio_quality"`
   ClassificationId         int    `json:"classification_id"`
   ContentId                string `json:"content_id"`
   ContentType              string `json:"content_type"`
   DeviceIdentifier         string `json:"device_identifier"`
   DeviceSerial             string `json:"device_serial"`
   DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
   Player                   string `json:"player"`
   SubtitleLanguage         string `json:"subtitle_language"`
   VideoType                string `json:"video_type"`
}

func (a Address) video(quality string) OnDemand {
   var v OnDemand
   v.AudioLanguage = "ENG"
   v.AudioQuality = "2.0"
   v.ContentType = "movies"
   v.DeviceSerial = "!"
   v.Player = "atvui40:DASH-CENC:WVM"
   v.SubtitleLanguage = "MIS"
   v.VideoType = "stream"
   v.DeviceIdentifier = "atvui40"
   v.ClassificationId = a.classification_id
   v.ContentId = a.content_id
   v.DeviceStreamVideoQuality = quality
   return v
}
