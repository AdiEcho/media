package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

// geo block
func (o on_demand) stream() (gizmo_stream, error) {
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   return io.ReadAll(res.Body)
}

type gizmo_stream []byte

func (g gizmo_stream) info() (*stream_info, error) {
   var s struct {
      Data struct {
         StreamInfos []stream_info `json:"stream_infos"`
      }
   }
   err := json.Unmarshal(g, &s)
   if err != nil {
      return nil, err
   }
   return &s.Data.StreamInfos[0], nil
}

type stream_info struct {
   LicenseUrl string `json:"license_url"`
   URL string
   VideoQuality string `json:"video_quality"`
}

func (w web_address) hd() on_demand {
   return w.video("HD")
}

func (w web_address) fhd() on_demand {
   return w.video("FHD")
}

func (s stream_info) RequestUrl() (string, bool) {
   return s.LicenseUrl, true
}

func (stream_info) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (stream_info) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (stream_info) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type on_demand struct {
   AudioLanguage string `json:"audio_language"`
   AudioQuality string `json:"audio_quality"`
   ClassificationId int `json:"classification_id"`
   ContentId string `json:"content_id"`
   ContentType string `json:"content_type"`
   DeviceIdentifier string `json:"device_identifier"`
   DeviceSerial string `json:"device_serial"`
   DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
   Player string `json:"player"`
   SubtitleLanguage string `json:"subtitle_language"`
   VideoType string `json:"video_type"`
}

func (w web_address) video(quality string) on_demand {
   var v on_demand
   v.AudioLanguage = "ENG"
   v.AudioQuality = "2.0"
   v.ContentType = "movies"
   v.DeviceSerial = "!"
   v.Player = "atvui40:DASH-CENC:WVM"
   v.SubtitleLanguage = "MIS"
   v.VideoType = "stream"
   v.DeviceIdentifier = "atvui40"
   v.ClassificationId = w.classification_id
   v.ContentId = w.content_id
   v.DeviceStreamVideoQuality = quality
   return v
}
