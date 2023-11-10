package hulu

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type deep_link struct {
   EAB_ID string
}

func (a authenticate) deeplink(id string) (*deep_link, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {id},
      "namespace": {"entity"},
      "user_token": {a.Data.User_Token},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   link := new(deep_link)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}
