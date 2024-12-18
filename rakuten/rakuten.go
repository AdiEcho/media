package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

func (s *StreamInfo) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      s.LicenseUrl, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type StreamInfo struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

// geo block
func (o *OnDemand) Info() (*StreamInfo, error) {
   data, err := json.Marshal(o)
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://gizmo.rakuten.tv/v3/avod/streamings",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var value struct {
      Data struct {
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.StreamInfos[0], nil
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

func (*GizmoMovie) Show() string {
   return ""
}

func (g *GizmoMovie) Title() string {
   return g.Data.Title
}

func (*GizmoMovie) Season() int {
   return 0
}

func (*GizmoMovie) Episode() int {
   return 0
}

type GizmoMovie struct {
   Data struct {
      Title string
      Year  int
   }
}

func (g *GizmoMovie) Year() int {
   return g.Data.Year
}
