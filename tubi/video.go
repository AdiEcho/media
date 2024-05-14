package tubi

import (
   "errors"
   "net/http"
   "slices"
   "strconv"
   "strings"
)

func (c Content) Video() (*VideoResource, error) {
   if len(c.V.VideoResources) == 0 {
      return nil, errors.New(".video_resources")
   }
   slices.SortFunc(c.V.VideoResources, func(a, b VideoResource) int {
      return int(b.Resolution - a.Resolution)
   })
   return &c.V.VideoResources[0], nil
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

func (VideoResource) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (VideoResource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (VideoResource) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type VideoResource struct {
   LicenseServer *struct {
      URL string
   } `json:"license_server"`
   Manifest struct {
      URL string
   }
   Resolution Resolution
   Type string
}

func (v VideoResource) RequestUrl() (string, bool) {
   if v := v.LicenseServer; v != nil {
      return v.URL, true
   }
   return "", false
}
