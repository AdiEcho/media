package stan

import (
   "net/http"
   "net/url"
)

func (a app_session) streams() (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://api.stan.com.au/concurrency/v1/streams", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "captions":[]string{"ttml"},
      "clientId":[]string{"6a25764ada16ddca"},
      "drm":[]string{"widevine"},
      "format":[]string{"dash"},
      "jwToken":[]string{a.JwToken},
      "manufacturer":[]string{"Android"},
      "model":[]string{"unknown"},
      "os":[]string{"Android"},
      "programId":[]string{"1768588"},
      "quality":[]string{"sd"},
      "sdk":[]string{"23"},
      "stanName":[]string{"Stan-Android"},
      "stanVersion":[]string{"4.31.1.50929"},
      "type":[]string{"mobile"},
      "videoCodec":[]string{"h264"},
   }.Encode()
   return http.DefaultClient.Do(req)
}
