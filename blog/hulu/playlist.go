package hulu

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

type playlist_request struct {
   Content_EAB_ID   string `json:"content_eab_id"`
   Deejay_Device_ID int    `json:"deejay_device_id"`
   Playback       struct {
      Audio struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values        []struct {
               Type string `json:"type"`
            } `json:"values"`
         } `json:"codecs"`
      } `json:"audio"`
      DRM struct {
         Selection_Mode string `json:"selection_mode"`
         Values        []struct {
            Security_Level string `json:"security_level"`
            Type          string `json:"type"`
            Version       string `json:"version"`
         } `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         Selection_Mode string `json:"selection_mode"`
         Values        []struct {
            Encryption struct {
               Mode string `json:"mode"`
               Type string `json:"type"`
            } `json:"encryption"`
            Type string `json:"type"`
         } `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
      Video   struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values        []struct {
               Level   string `json:"level"`
               Profile string `json:"profile"`
               Type    string `json:"type"`
            } `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
   } `json:"playback"`
   Token          string `json:"token"`
   Unencrypted    bool   `json:"unencrypted"`
   Version        int    `json:"version"`
}

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Scheme = "https"
   req.URL.Host = "play.hulu.com"
   req.URL.Path = "/v6/playlist"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/json"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
