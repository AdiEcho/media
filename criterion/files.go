package criterion

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

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

func (a AuthToken) Files(item *EmbedItem) (VideoFiles, error) {
   req, err := http.NewRequest("", item.Links.Files.Href, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer "+a.AccessToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var files VideoFiles
   err = json.NewDecoder(resp.Body).Decode(&files)
   if err != nil {
      return nil, err
   }
   return files, nil
}

func (v VideoFiles) Dash() (*VideoFile, bool) {
   for _, file := range v {
      if file.Method == "dash" {
         return &file, true
      }
   }
   return nil, false
}
