package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

type program_stream struct {
   Media struct {
      DRM *struct {
         CustomData string
         KeyId string
      }
      VideoUrl string
   }
}

func (a app_session) stream(id int) (*program_stream, error) {
   req, err := http.NewRequest(
      "GET", "https://api.stan.com.au/concurrency/v1/streams", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("x-forwarded-for", "1.128.0.0")
   req.URL.RawQuery = url.Values{
      "drm": {"widevine"}, // need for .Media.DRM
      "format": {"dash"}, // 404 otherwise
      "jwToken": {a.JwToken},
      "programId": {strconv.Itoa(id)},
      "quality": {"auto"}, // note `high` or `ultra` should work too
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stream := new(program_stream)
   if err := json.NewDecoder(res.Body).Decode(stream); err != nil {
      return nil, err
   }
   return stream, nil
}

func (program_stream) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (p program_stream) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("dt-custom-data", p.Media.DRM.CustomData)
   return head, nil
}

func (program_stream) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (program_stream) ResponseBody(b []byte) ([]byte, error) {
   var s struct {
      License []byte
   }
   err := json.Unmarshal(b, &s)
   if err != nil {
      return nil, err
   }
   return s.License, nil
}

// `akamaized.net` geo blocks, so change the host. note `aws.stan.video`
// should work too
func (p program_stream) StanVideo() (*url.URL, error) {
   video, err := url.Parse(p.Media.VideoUrl)
   if err != nil {
      return nil, err
   }
   video.Host = "gec.stan.video"
   return video, nil
}
