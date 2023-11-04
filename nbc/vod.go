package nbc

import (
   "net/http"
   "net/url"
   "strconv"
)

const mpx_account_id = 2410887629

func VOD(mpx_guid int64) (*http.Response, error) {
   req, err := http.NewRequest("GET", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, mpx_account_id, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, mpx_guid, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {"Full Episode"},
   }.Encode()
   return http.DefaultClient.Do(req)
}
