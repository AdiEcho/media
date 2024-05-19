package criterion

import (
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

type video_file struct {
   DrmAuthorizationToken string `json:"drm_authorization_token"`
   Links struct {
      Source struct {
         Href string
      }
   } `json:"_links"`
   Method string
}

func (video_file) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (v video_file) RequestUrl() (string, bool) {
   b := []byte("https://drm.vhx.com/v2/widevine?token=")
   b = append(b, v.DrmAuthorizationToken...)
   return string(b), true
}

func (video_file) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (video_file) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

type video_files []video_file

func (v video_files) dash() (*video_file, bool) {
   for _, file := range v {
      if file.Method == "dash" {
         return &file, true
      }
   }
   return nil, false
}

func (a AuthToken) files(item *embed_item) (video_files, error) {
   address := func() string {
      b := []byte("https://api.vhx.com/videos/")
      b = strconv.AppendInt(b, item.ID, 10)
      b = append(b, "/files"...)
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
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var files video_files
   err = json.NewDecoder(res.Body).Decode(&files)
   if err != nil {
      return nil, err
   }
   return files, nil
}
