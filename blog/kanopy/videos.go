package kanopy

import (
   "net/http"
   "strconv"
)

func (w *web_token) videos(id int) (*http.Response, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/videos/" + strconv.Itoa(id)
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   return http.DefaultClient.Do(req)
}
