package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
)

func one() {
   username := os.Getenv("spotify_username")
   if username == "" {
      panic("spotify_username")
   }
   password := os.Getenv("spotify_password")
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
      }},
      protobuf.Field{Number: 101, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(username)},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes(password)},
      }},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "login5.spotify.com"
   req.URL.Path = "/v3/login"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Encode()))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   body = nil
   if err := body.Consume(data); err != nil {
      panic(err)
   }
   res.Write(os.Stdout)
   fmt.Printf("%#v\n", body)
}
