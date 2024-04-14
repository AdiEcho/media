package tubi

import (
   "net/http"
   "slices"
   "strconv"
   "strings"
)

func (c Content) Video() VideoResource {
   slices.SortFunc(c.Video_Resources, func(a, b VideoResource) int {
      return int(b.Resolution - a.Resolution)
   })
   return c.Video_Resources[0]
}

type Resolution int

func (r *Resolution) UnmarshalText(text []byte) error {
   a := strings.TrimPrefix(string(text), "VIDEO_RESOLUTION_")
   i, err := strconv.Atoi(strings.TrimSuffix(a, "P"))
   if err != nil {
      return err
   }
   *r = Resolution(i)
   return nil
}

type VideoResource struct {
   License_Server struct {
      URL string
   }
   Manifest struct {
      URL string
   }
   Resolution Resolution
}

func (VideoResource) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (VideoResource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (v VideoResource) RequestUrl() (string, bool) {
   return v.License_Server.URL, true
}

func (VideoResource) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}
