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

// wikipedia.org/wiki/Hashcash
func hashcash(login_context, message []byte) []byte {
   rand := func() uint64 {
      sum := sha1.Sum(login_context)
      return binary.BigEndian.Uint64(sum[len(sum)-8:])
   }()
   var counter uint64
   for {
      var b []byte
      b = binary.BigEndian.AppendUint64(b, rand)
      b = binary.BigEndian.AppendUint64(b, counter)
      zero_bits := func() int {
         sum := sha1.Sum(append(message, b...))
         x := binary.BigEndian.Uint16(sum[sha1.Size-2:])
         return bits.TrailingZeros16(x)
      }()
      if zero_bits >= 10 {
         return b
      }
      rand++
      counter++
   }
}

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
      m.AddBytes(1, []byte(client_id))
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
