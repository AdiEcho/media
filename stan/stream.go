package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

func (a AppSession) Stream(id int64) (*ProgramStream, error) {
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
      "programId": {strconv.FormatInt(id, 10)},
      "quality": {"auto"}, // note `high` or `ultra` should work too
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stream := new(ProgramStream)
   if err := json.NewDecoder(res.Body).Decode(stream); err != nil {
      return nil, err
   }
   return stream, nil
}

func (ProgramStream) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (p ProgramStream) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("dt-custom-data", p.Media.DRM.CustomData)
   return head, nil
}

func (ProgramStream) RequestUrl() (string, bool) {
   return "https://lic.drmtoday.com/license-proxy-widevine/cenc/", true
}

func (ProgramStream) ResponseBody(b []byte) ([]byte, error) {
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
func (p ProgramStream) StanVideo() (*url.URL, error) {
   video, err := url.Parse(p.Media.VideoUrl)
   if err != nil {
      return nil, err
   }
   video.Host = "gec.stan.video"
   return video, nil
}
type ProgramStream struct {
   Media struct {
      DRM *struct {
         CustomData string
         KeyId string
      }
      VideoUrl string
   }
}
