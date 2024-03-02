package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

func (r login_response) solve_hash_cash_challenge() bool {
   return false
}

// protobuf.Message{
//    protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
//       protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
//          protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
//             protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xd6M˔?\xc6\x13\x18\x95\xa7\xab\xd3\xf8\b\xaea")},
//             protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(10)},
//          }},
//       }},
//    }},
//    protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("\x03\x003Oub\t\xc9\x1684\x84\xa7\x8c\x1d]8db\xd3\x03\x1e\xf7\x1f\xfb\x05\xe0S\x0f\xc3\x03\x99\xe8\xe2\x94\xda\xea]ba\xf6wÆ\x8aо\xe1p\x13\xfa\x88\xe9\xbe2%\xa6\xf6\x7f\xeb˶qE\xdb\xdb\x11\x05\x13\xef\xe7\x06\xf3F\xb8\x90\xd1x\xd8\xd6\x1d\n\xba\xd2\n\x8c\xfaȗ8\x89\xfdC\xe4,=\xb66;z\xe3\xc5E^\x8c\xef\xef\xd6JF\x1a^*\x99\xf5\x10=m\xca\xf0\xc8%\xd3[6\xc5s}\x85+\xbfZ7j\x1d\x9e'\xf6!?\xf0\x1d\xd9\xe1D@\xe2\xa0n\xb6]`\xbe@Y\xfe\xa8>\xc8\xfd\x05Lظ\xd3\x13\xba&\xc4\\\x15\xef\xf647}\a\xc7|S\xed\xf6\xe3{\x81\xf8p:\xe0\x12rD%Jѧ(\xeb\x02ul;\x8a\x8d\xb9b\xc4VX|\x1c>\xd1P\xe9\xbeY\xf8\xa9\xca\x1d\xf8((")},
// }
type login_response struct {
   m protobuf.Message
}

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
