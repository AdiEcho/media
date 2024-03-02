package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

type login_request struct {
   m protobuf.Message
}

func (r *login_request) New(username, password string) bool {
   if username == "" {
      return false
   }
   if password == "" {
      return false
   }
   r.m = protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
      }},
      protobuf.Field{Number: 101, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(username)},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes(password)},
      }},
   }
   return true
}

type login_response struct {
   m protobuf.Message
}

func (r login_request) login() (*login_response, error) {
   req, err := http.NewRequest(
      "POST", "https://login5.spotify.com/v3/login",
      bytes.NewReader(r.m.Encode()),
   )
   if err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var login login_response
   if err := login.m.Consume(data); err != nil {
      return nil, err
   }
   return &login, nil
}
