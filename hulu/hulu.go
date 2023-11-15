package hulu

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
)

type Details struct {
   Items []struct {
      Series_Name string
      Episode_Name string
   }
}

func (a Authenticate) Details(d Deep_Link) (*Details, error) {
   body, err := func() ([]byte, error) {
      m := map[string][]string{
         "eabs": {d.EAB_ID},
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://guide.hulu.com/guide/details?user_token=" + a.Data.User_Token,
      "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   detail := new(Details)
   if err := json.NewDecoder(res.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}

type Deep_Link struct {
   EAB_ID string
}

func (a Authenticate) Deep_Link(id string) (*Deep_Link, error) {
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
   link := new(Deep_Link)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}
