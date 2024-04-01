package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type program_streams struct {
   Media struct {
      DRM struct {
         CustomData string
         KeyId string
      }
      VideoUrl string
   }
}

func (a app_session) program() (*program_streams, error) {
   req, err := http.NewRequest(
      "GET", "https://api.stan.com.au/concurrency/v1/streams", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header["x-forwarded-for"] = []string{"1.128.0.0"}
   req.URL.RawQuery = url.Values{
      "drm": {"widevine"},
      "format": {"dash"},
      "jwToken": {a.JwToken},
      "programId": {"1768588"},
      "quality": {"auto"}, // note `high` or `ultra` should work too
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   program := new(program_streams)
   if err := json.NewDecoder(res.Body).Decode(program); err != nil {
      return nil, err
   }
   return program, nil
}

// `akamaized.net` geo blocks, so change the host. note `aws.stan.video`
// should work too
func (p program_streams) StanVideo() (*url.URL, error) {
   video, err := url.Parse(p.Media.VideoUrl)
   if err != nil {
      return nil, err
   }
   video.Host = "gec.stan.video"
   return video, nil
}

func (program_streams) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (program_streams) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (p program_streams) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("dt-custom-data", p.Media.DRM.CustomData)
   return head, nil
}

func (program_streams) ResponseBody(b []byte) ([]byte, error) {
   var s struct {
      License []byte
   }
   err := json.Unmarshal(b, &s)
   if err != nil {
      return nil, err
   }
   return s.License, nil
}
