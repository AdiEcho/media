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

func Two() {
   username := os.Getenv("spotify_username")
   if username == "" {
      panic("spotify_username")
   }
   password := os.Getenv("spotify_password")
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("58cebdd226ac462a")},
      }},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("\x03\x00\xa4\x12\xf2\xcb9\x1e\xc9\x0e\v z\xfe/\xa9\xf9\x9a=\xa2\x1a\f\xb6\xab\x9e=\xef\xa6q<\xa2恨a\x05\xa6kǪ\xa0\xcc\xd99]+ثHX\xca\xe8h\x85,\x02\xe4I\x05i\xb5\xc9/\xea#yզ\x1e܈UG\x18\xa5\\\xe4\xf2\xde\xea.\xfd\xf3\x1a\xa7\xed\x06N\xea\xb8\x026|\x17\x06\xae<)_R\x1e\xa0\xbebfG\x94\xf5i\xb6\x91\x00\x88ر\x90G{\xf4\xe0d\xb6\x11\x82\x16\xb5\xc0\n\x81HZ\xd6g\xd3K\x96\xb0:\x0f\x8eH^\xba\n\xc7.3UJ\xb6\xc88\x02\xd1.\xf5\x8b\x94ќb\x89\a\xd3DI:Fur\x89\xf6\xa4\xddtL\x1a\xfbso\xe6\x11\xc6μo\xb1\xb7\x99\x8a\x1b\xae\x10[\xf7\xb7\x19=\xacU\xb0\x19\x01\x1b\x05&\xbaZ\x02r\xa6\xab\xff\xea\x1b\x19\xdb\ra\xd8R\xb9'{\x12*]\xe2\xa7(\\\x06x#\x8a@}\xe0\x98_\x03-e\xbe\xec\xbc:\xf1\xc4\x12\x92\xe5[\xe7\xacd\xd6\x10H@춲\xe8\xf5L\xf5\xf4\xeeC\xd1\x02\xa6\xbf\x8bc\xbf\x8b\x8c\xe6\xda")},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xa84\xe7\xf9\x8aֵ\xbb\x00\x00\x00\x00\x00\x00\x006")},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Message{
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(14400)},
               }},
            }},
         }},
      }},
      protobuf.Field{Number: 101, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(username)},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes(password)},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")},
      }},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Cache-Control"] = []string{"no-cache, no-store, max-age=0"}
   req.Header["Client-Token"] = []string{"AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q=="}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"Spotify/8.9.18.512 Android/23 (Android SDK built for x86)"}
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
