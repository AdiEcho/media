package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/url"
)

func (login_ok) metadata() (*http.Response, error) {
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("US")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("free")},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("\x0e\xfdN\x9d\x9c\xd8.y\x95\xd1%\xb9\xa7\x01,\xf8")},
      }},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("spotify:track:1oaaSrDJimABpOdCEbw2DJ")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(5)},
         }},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(10)},
         }},
      }},
   }
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/extended-metadata/v0/extended-metadata"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Encode()))
   req.Header["Authorization"] = []string{"Bearer BQCkkXlvEzT-iTS4rlLwOnnAzmyxcuz7yI19Joys5qvLZxwB0XCm8bea7ikhOoioxprBD8jGa0gqnBq1wSIUXbi6Yt9iB-uZYRv5Ogwu6Ccq_59CfHlB6x8dzHeFxuvGVvQCdCQ7RMZfZ3aucXPXNNMnt_Pm8hp1dNLGeb92CKWSIf7f6UziCrBVTfJap2f0j_uHbjZamT3DKve-xhj0ViqHA30WPY6EZFhs6pzAAPmBp4hjNmheQvwMU9GWhKjvxVlJvbRV994gWlg01krDWis4CC7CsEVKOVRYBCIkg3H5vl5ymO2dNFuVvFQSCmUuWYPqx350UmulKbObUvzz"}
   return http.DefaultClient.Do(req)
}
