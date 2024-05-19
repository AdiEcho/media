package criterion

import "net/http"

func (a AuthToken) video(slug string) (*http.Response, error) {
   req, err := http.NewRequest("", "https://api.vhx.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/videos/" + slug
   req.URL.RawQuery = "url=" + slug
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
   return http.DefaultClient.Do(req)
}
