package ctv

import (
   "net/http"
   "net/url"
   "os"
)

func Manifest() {
   var req http.Request
   req.URL = new(url.URL)
   req.URL.Scheme = "https"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL.Host = "capi.9c9media.com"
   req.URL.Path = "/destinations/ctvmovies_hub/platforms/desktop/playback/contents/1417780/contentPackages/2852723/manifest.mpd"
   req.Header = make(http.Header)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
