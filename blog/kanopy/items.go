package kanopy

import (
   "net/http"
   "strconv"
)

func (w *web_token) items(video_id int64) (*http.Response, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/kapi/videos/")
      b = strconv.AppendInt(b, video_id, 10)
      b = append(b, "/items"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   return http.DefaultClient.Do(req)
}
