package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

// "file_id": "f682d2a95d0e14eeef4f40b60fddde56bc6721c7",
// "format": "OGG_VORBIS_96"
// curl -o OGG_VORBIS_320 https://audio4-fa.scdn.co/audio/98b53c239db400b0a016145700de431f68d28f54?1710302675_Jq6zxYruifAH49PewVa2zM6T3xzhR41E7EWxkKpFXxs=
// curl -o OGG_VORBIS_160 https://audio4-fa.scdn.co/audio/6a5f12fa51f2c1e284af707a99f3ca8696f7f62f?1710302676_AF10C-b4HwVkXT7zUu2Q6ttpWANUJzkvAHq4fT1p8RE=
// curl -o OGG_VORBIS_96 https://audio4-fa.scdn.co/audio/f682d2a95d0e14eeef4f40b60fddde56bc6721c7?1710302677_ftEQKNOL1OYk52ebKBJSipH6sQGaAvcFIkhnvynHIfs=
func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/playplay/v1/key/f682d2a95d0e14eeef4f40b60fddde56bc6721c7"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Encode()))
   req.Header["Authorization"] = []string{"Bearer BQCHAsDof0SnWGNOSmBjfxXr0cC_eqdLY8N_fL_54XlqESZOQ2qZ0q65s4TWOIMXDXF1siJgUnfZKAei44OSIuIXZovqA2go5byZ5MEFYshZYlUT3Jh2EKw2wIRyaXhjlaAxtqizrr4uF0izrNYPpCzk61Um1idTuJXwTVSMxkXrdWEN9RcNMh0nSZzCmhCkDtcvokHGSd3MSK7QiwFad_d80FaoG2b4PivFzfllNic6a6m4F4t22ztD9Ho70I8Tb30M5ewyS5aztG5BF29rZiNqNdvAoS5evDSLOuoz32mKMM_CNV7uGSyaySW-W0C9m9B1Knw19JKhDyH2NX6_"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if err := body.Consume(data); err != nil {
      panic(err)
   }
   fmt.Printf("%x\n", <-body.GetBytes(1))
}

var body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(2)},
   protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("\x01K\xe0K\xce^\xe6\xb3nl\xec0\xd8\xeb\x9a2")},
}
