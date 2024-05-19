package criterion

import (
   "net/http"
   "net/url"
   "strconv"
)

func (a AuthToken) files(item *embed_item) (*http.Response, error) {
   address := func() string {
      b := []byte("https://api.vhx.com/videos/")
      b = strconv.AppendInt(b, item.ID, 10)
      b = append(b, "/files"...)
      return string(b)
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "codec": {"h264"},
      "format": {"mpd"},
      "quality": {"adaptive"},
   }.Encode()
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
   return http.DefaultClient.Do(req)
}
