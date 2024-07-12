package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

// github.com/mitmproxy/mitmproxy/blob/main/mitmproxy/contentviews/protobuf.py
func (StreamInfo) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set("content-type", "application/x-protobuf")
   return h, nil
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
   var s struct {
      Data struct {
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data.StreamInfos[0], nil
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
