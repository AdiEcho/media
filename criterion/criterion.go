package criterion

import "net/http"

type VideoFile struct {
   DrmAuthorizationToken string `json:"drm_authorization_token"`
   Links                 struct {
      Source struct {
         Href string
      }
   } `json:"_links"`
   Method string
}

type VideoFiles []VideoFile

func (VideoFile) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (v VideoFile) RequestUrl() (string, bool) {
   b := []byte("https://drm.vhx.com/v2/widevine?token=")
   b = append(b, v.DrmAuthorizationToken...)
   return string(b), true
}

func (VideoFile) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (VideoFile) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (v VideoFiles) Dash() (*VideoFile, bool) {
   for _, file := range v {
      if file.Method == "dash" {
         return &file, true
      }
   }
   return nil, false
}

type EmbedItem struct {
   Links struct {
      Files struct {
         Href string
      }
   } `json:"_links"`
   Metadata struct {
      YearReleased int `json:"year_released"`
   }
   Name string
}

func (EmbedItem) Episode() int {
   return 0
}

func (EmbedItem) Season() int {
   return 0
}

func (EmbedItem) Show() string {
   return ""
}

func (e EmbedItem) Title() string {
   return e.Name
}

func (e EmbedItem) Year() int {
   return e.Metadata.YearReleased
}

const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"
