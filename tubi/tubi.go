package tubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (v *VideoContent) New(id int) error {
   req, err := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(id)},
      "deviceId":   {"!"},
      "platform":   {"android"},
      "video_resources[]": {
         "dash",
         "dash_widevine",
      },
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   v.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (v *VideoContent) Unmarshal() error {
   err := json.Unmarshal(v.Raw, v)
   if err != nil {
      return err
   }
   if len(v.VideoResources) == 0 {
      return errors.New(string(v.Raw))
   }
   v.set(nil)
   return nil
}

func (v *VideoContent) Get(id int) (*VideoContent, bool) {
   if v.Id == id {
      return v, true
   }
   for _, child := range v.Children {
      if content, ok := child.Get(id); ok {
         return content, true
      }
   }
   return nil, false
}

type VideoContent struct {
   Children       []*VideoContent
   DetailedType   string `json:"detailed_type"`
   EpisodeNumber  int    `json:"episode_number,string"`
   Id             int    `json:",string"`
   Raw            []byte `json:"-"`
   SeriesId       int    `json:"series_id,string"`
   Title          string
   VideoResources []VideoResource `json:"video_resources"`
   Year           int
   parent         *VideoContent
}

func (v *VideoContent) set(parent *VideoContent) {
   v.parent = parent
   for _, child := range v.Children {
      child.set(v)
   }
}

// VideoContent.Unmarshal checks the length
func (v *VideoContent) Video() VideoResource {
   a := v.VideoResources[0]
   for _, b := range v.VideoResources {
      if b.Resolution.Int64 > a.Resolution.Int64 {
         a = b
      }
   }
   return a
}

func (v *VideoContent) Episode() bool {
   return v.DetailedType == "episode"
}

func (r Resolution) MarshalText() ([]byte, error) {
   b := []byte("VIDEO_RESOLUTION_")
   b = strconv.AppendInt(b, r.Int64, 10)
   return append(b, 'P'), nil
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

func (VideoResource) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (VideoResource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (VideoResource) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (n Namer) Episode() int {
   return n.Content.EpisodeNumber
}

func (v *VideoResource) RequestUrl() (string, bool) {
   if v.LicenseServer != nil {
      return v.LicenseServer.Url, true
   }
   return "", false
}

func (n Namer) Year() int {
   return n.Content.Year
}

type Namer struct {
   Content *VideoContent
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
