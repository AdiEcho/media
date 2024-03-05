package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (o login_ok) metadata() (*http.Response, error) {
   token, ok := o.access_token()
   if !ok {
      return nil, errors.New("login_ok.access_token")
   }
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
   req, err := http.NewRequest(
      "POST", "https://guc3-spclient.spotify.com",
      bytes.NewReader(body.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/extended-metadata/v0/extended-metadata"
   req.Header.Set("authorization", "Bearer " + token)
   return http.DefaultClient.Do(req)
}
