package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "bytes"
   "154.pages.dev/protobuf"
   "fmt"
)

const video_ID = "oCjW6gdEDa4"

func main() {
   var req_body = protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 16, Type: 0,  Value: protobuf.Varint(3)},
            protobuf.Field{Number: 17, Type: 2,  Value: protobuf.Bytes("16.49.39")},
            protobuf.Field{Number: 19, Type: 2,  Value: protobuf.Bytes("12")},
            protobuf.Field{Number: 64, Type: 0,  Value: protobuf.Varint(32)},
         }},
      }},
      protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes(video_ID)},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Append(nil)))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/16.49.39(Linux; U; Android 12; en_US; sdk_gphone64_x86_64 Build/SE1B.220616.007) gzip"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if strings.Contains(string(data), "This video requires payment to watch") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
