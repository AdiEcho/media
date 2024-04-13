package tubi

import "net/http"

func (v video_resource) RequestUrl() (string, bool) {
   return v.License_Server.URL, true
}

type video_resource struct {
   License_Server struct {
      URL string
   }
   Manifest struct {
      URL string
   }
   Resolution string
}

func (video_resource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (video_resource) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (video_resource) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

func (c content) resolution_720p() (*video_resource, bool) {
   for _, video := range c.Video_Resources {
      if video.Resolution == "VIDEO_RESOLUTION_720P" {
         return &video, true
      }
   }
   return nil, false
}
