package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

var classification = map[string]int{
   "fr": 23,
   "se": 282,
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

func (o *on_demand) New(class int, content_id string) {
   o.AudioLanguage = "ENG"
   o.AudioQuality = "2.0"
   o.ClassificationId = class
   o.ContentId = content_id
   o.ContentType = "movies"
   o.DeviceIdentifier = "atvui40"
   o.DeviceSerial = "!"
   o.DeviceStreamVideoQuality = "FHD"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
}

func (o on_demand) stream() (*stream_info, error) {
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
         StreamInfos []stream_info `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data.StreamInfos[0], nil
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

type stream_info struct {
   LicenseUrl string `json:"license_url"`
   URL string
   VideoQuality string `json:"video_quality"`
}

func (stream_info) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
