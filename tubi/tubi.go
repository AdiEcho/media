package tubi

import (
   "net/http"
   "strconv"
   "strings"
)

type Namer struct {
   Content *VideoContent
}

func (n Namer) Episode() int {
   return n.Content.EpisodeNumber
}

func (n Namer) Season() int {
   if n.Content.parent != nil {
      return n.Content.parent.Id
   }
   return 0
}

func (n Namer) Show() string {
   if v := n.Content.parent; v != nil {
      return v.parent.Title
   }
   return ""
}

// S01:E03 - Hell Hath No Fury
func (n Namer) Title() string {
   if _, v, ok := strings.Cut(n.Content.Title, " - "); ok {
      return v
   }
   return n.Content.Title
}

func (n Namer) Year() int {
   return n.Content.Year
}

type Resolution struct {
   Int64 int64
}

func (r *Resolution) UnmarshalText(text []byte) error {
   s := string(text)
   s = strings.TrimPrefix(s, "VIDEO_RESOLUTION_")
   s = strings.TrimSuffix(s, "P")
   var err error
   r.Int64, err = strconv.ParseInt(s, 10, 64)
   if err != nil {
      return err
   }
   return nil
}

type VideoResource struct {
   LicenseServer *struct {
      Url string
   } `json:"license_server"`
   Manifest struct {
      Url string
   }
   Resolution Resolution
   Type       string
}

func (v *VideoResource) RequestUrl() (string, bool) {
   if v.LicenseServer != nil {
      return v.LicenseServer.Url, true
   }
   return "", false
}

func (r Resolution) MarshalText() ([]byte, error) {
   b := []byte("VIDEO_RESOLUTION_")
   b = strconv.AppendInt(b, r.Int64, 10)
   return append(b, 'P'), nil
}

func (*VideoResource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (*VideoResource) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (*VideoResource) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}
