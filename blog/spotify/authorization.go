package spotify

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

func (r login_response) challenge(
   username, password string,
) (protobuf.Message, error) {
   login_context, ok := r.m.GetBytes(5)
   if !ok {
      return nil, errors.New("login_context")
   }
   prefix, ok := func() ([]byte, bool) {
      m, _ := r.m.Get(3)
      m, _ = m.Get(1)
      m, _ = m.Get(1)
      return m.GetBytes(1)
   }()
   if !ok {
      return nil, errors.New("prefix")
   }
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.AddBytes(1, []byte(android_client_id))
   })
   m.AddBytes(2, login_context)
   m.Add(3, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add(1, func(m *protobuf.Message) {
            m.AddBytes(1, hashcash(login_context, prefix))
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
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["User-Agent"] = []string{"Symfony HttpClient (Curl)"}
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   m = nil
   if err := m.Consume(data); err != nil {
      return nil, err
   }
   return m, nil
}
