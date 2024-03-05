package android

import (
   "154.pages.dev/protobuf"
   "bytes"
   "crypto/sha1"
   "encoding/binary"
   "errors"
   "io"
   "math/bits"
   "net/http"
)

func (h login_response) ok(username, password string) (*LoginOk, error) {
   login_context, ok := h.login_context()
   if !ok {
      return nil, errors.New("login_response.login_context")
   }
   prefix, ok := h.prefix()
   if !ok {
      return nil, errors.New("login_response.prefix")
   }
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.AddBytes(1, []byte(android_client_id))
   })
   m.AddBytes(2, login_context)
   m.Add(3, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add(1, func(m *protobuf.Message) {
            m.AddBytes(1, solve_hash_cash(login_context, prefix, 10))
         })
      })
   })
   m.Add(101, func(m *protobuf.Message) {
      m.AddBytes(1, []byte(username))
      m.AddBytes(2, []byte(password))
   })
   req, err := http.NewRequest(
      "POST", "https://login5.spotify.com/v3/login", bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var login LoginOk
   login.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &login, nil
}

// github.com/librespot-org/librespot/blob/dev/protocol/proto/spotify/login5/v3/login5.proto
type LoginOk struct {
   Data []byte
   m protobuf.Message
}

func (o *LoginOk) Consume() error {
   return o.m.Consume(o.Data)
}

func (o LoginOk) AccessToken() (string, bool) {
   if v, ok := o.m.Get(1); ok { // LoginOk ok
      if v, ok := v.GetBytes(2); ok { // string access_token
         return string(v), true
      }
   }
   return "", false
}
// github.com/librespot-org/librespot/blob/dev/core/src/spclient.rs
func solve_hash_cash(login_context, prefix []byte, length int) []byte {
   sum := sha1.Sum(login_context)
   var counter uint64
   target := binary.BigEndian.Uint64(sum[12:])
   for {
      suffix := func() []byte {
         b := binary.BigEndian.AppendUint64(nil, target)
         return binary.BigEndian.AppendUint64(b, counter)
      }()
      sum := sha1.Sum(append(prefix, suffix...))
      if bits.TrailingZeros64(binary.BigEndian.Uint64(sum[12:])) >= length {
         return suffix
      }
      counter++
      target++
   }
}

