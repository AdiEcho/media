package rakuten

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

var classification = map[string]int{
   "fr": 23,
   "se": 282,
}

func streamings(class int, content_id string) (*http.Response, error) {
   body, err := json.Marshal(map[string]string{
      "audio_language": "ENG",
      "audio_quality": "2.0",
      "classification_id": strconv.Itoa(class),
      "content_id": content_id,
      "content_type": "movies",
      "device_serial": "!",
      "player": "atvui40:DASH-CENC:WVM",
      "subtitle_language": "MIS",
      "video_type": "stream",
   })
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
   req.URL.RawQuery = url.Values{
      "device_identifier": {"atvui40"},
      "device_stream_video_quality": {"FHD"},
   }.Encode()
   req.Header.Set("content-type", "application/json")
   return http.DefaultClient.Do(req)
}
