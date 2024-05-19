package criterion

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (a auth_token) delivery(item *embed_item) (*video_delivery, error) {
   address := func() string {
      b := []byte("https://api.vhx.com/v2/sites/59054/videos/")
      b = strconv.AppendInt(b, item.ID, 10)
      b = append(b, "/delivery"...)
      return string(b)
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   video := new(video_delivery)
   err = json.NewDecoder(res.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

type video_delivery struct {
   Streams []video_stream
}

func (v video_delivery) dash() (*video_stream, bool) {
   for _, stream := range v.Streams {
      if stream.MaxHeight == nil {
         if stream.Method == "dash" {
            return &stream, true
         }
      }
   }
   return nil, false
}

type video_stream struct {
   DRM struct {
      Schemes struct {
         Widevine struct {
            LicenseUrl string `json:"license_url"`
         }
      }
   }
   MaxHeight *int `json:"max_height"`
   Method string
   URL string
}

func (video_stream) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (v video_stream) RequestUrl() (string, bool) {
   return v.DRM.Schemes.Widevine.LicenseUrl, true
}

func (video_stream) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (video_stream) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}
