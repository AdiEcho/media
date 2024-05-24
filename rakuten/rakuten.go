package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (s *StreamInfo) Unmarshal(text []byte) error {
   return json.Unmarshal(text, s)
}

func (s StreamInfo) Marshal() ([]byte, error) {
   return json.MarshalIndent(s, "", " ")
}

type StreamInfo struct {
   LicenseUrl   string `json:"license_url"`
   URL          string
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
   var s struct {
      Data struct {
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data.StreamInfos[0], nil
}

func (w WebAddress) FHD() OnDemand {
   return w.video("FHD")
}

func (w WebAddress) hd() OnDemand {
   return w.video("HD")
}

func (s StreamInfo) RequestUrl() (string, bool) {
   return s.LicenseUrl, true
}

func (StreamInfo) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (StreamInfo) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
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

func (w WebAddress) video(quality string) OnDemand {
   var v OnDemand
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

