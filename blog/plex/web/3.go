package plex

import (
   "net/http"
   "net/url"
   "os"
)

func play_queues() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play.provider.plex.tv"
   req.URL.Path = "/playQueues"
   req.URL.Scheme = "https"
   req.Header["Accept"] = []string{"application/json"}
   req.Header["X-Plex-Token"] = []string{"fc1WPqnLdmq3J4Axt5pn"}
   val := make(url.Values)
   val["uri"] = []string{"provider://tv.plex.provider.vod/library/metadata/5d776d15f617c90020185cc6"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
