package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "errors"
   "net/http"
   "net/url"
)

func metadata(login android.LoginOk, track string) (*http.Response, error) {
   token, ok := login.AccessToken()
   if !ok {
      return nil, errors.New("android.LoginOk.AccessToken")
   }
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "spclient.wg.spotify.com"
   req.URL.Path = "/metadata/4/track/" + track
   val := make(url.Values)
   val["market"] = []string{"from_token"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Authorization"] = []string{"Bearer " + token}
   return http.DefaultClient.Do(&req)
}
