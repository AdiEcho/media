package sbs

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type video_stream struct {
   StreamProviders []struct {
      ContentUrl string
   }
}

func (a auth_native) video() (*video_stream, error) {
   req, err := http.NewRequest(
      "GET", "https://www.sbs.com.au/api/v3/video_stream", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + a.User_State.SessionToken)
   req.URL.RawQuery = url.Values{
      "context": {"odwebsite"},
      //"context": {"tv"},
      "id": {"2229616195516"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   video := new(video_stream)
   err = json.NewDecoder(res.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}
