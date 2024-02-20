package mubi

import (
   "net/http"
   "net/url"
   "os"
)

// https://mubi.com/en/us/films/passages-2022
// https://mubi.com/en/films/325455/player

const bearer = "79df72c589f824c61245ba11ff0a4c63fb13a1"

func Three() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.mubi.com"
   req.URL.Path = "/v3/films/187/viewing/secure_url"
   req.URL.Scheme = "https"
   req.Header["Client"] = []string{"web"}
   req.Header["Client-Country"] = []string{"GR"}
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
