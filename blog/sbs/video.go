package sbs

import (
   "net/http"
   "net/url"
)

func (a auth_native) video_stream() (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://www.sbs.com.au/api/v3/video_stream", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + a.User_State.SessionToken)
   req.URL.RawQuery = url.Values{
      "context": {"odwebsite"},
      "id": {"2229616195516"},
   }.Encode()
   return http.DefaultClient.Do(req)
}
