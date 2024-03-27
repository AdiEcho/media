package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
)

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

func (a app_session) program() (*program_streams, error) {
   req, err := http.NewRequest(
      "GET", "https://api.stan.com.au/concurrency/v1/streams", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header["x-forwarded-for"] = []string{"1.128.0.0"}
   req.URL.RawQuery = url.Values{
      "format":[]string{"dash"},
      "jwToken":[]string{a.JwToken},
      "programId":[]string{"1768588"},
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

type program_streams struct {
   Media struct {
      VideoUrl string
   }
}
