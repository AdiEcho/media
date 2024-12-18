package tubi

import (
   "bytes"
   "io"
   "net/http"
   "strconv"
   "strings"
)

func (v *VideoResource) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      v.LicenseServer.Url, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
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

type Resolution struct {
   Data int64
}

func (r *Resolution) UnmarshalText(text []byte) error {
   s := string(text)
   s = strings.TrimPrefix(s, "VIDEO_RESOLUTION_")
   s = strings.TrimSuffix(s, "P")
   var err error
   r.Data, err = strconv.ParseInt(s, 10, 64)
   if err != nil {
      return err
   }
   return nil
}

func (r Resolution) MarshalText() ([]byte, error) {
   b := []byte("VIDEO_RESOLUTION_")
   b = strconv.AppendInt(b, r.Data, 10)
   return append(b, 'P'), nil
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

type Namer struct {
   Content *VideoContent
}

func (n Namer) Episode() int {
   return n.Content.EpisodeNumber
}

func (n Namer) Year() int {
   return n.Content.Year
}

func (n Namer) Season() int {
   if n.Content.parent != nil {
      return n.Content.parent.Id
   }
   return 0
}
