package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "crypto/sha1"
   "errors"
   "io"
   "net/http"
)

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

func (r login_response) solve_hash_cash_challenge() (
   []byte, int, time.Duration, error,
) {
   hash_cash_challenge, _ := func() (m protobuf.Message, ok bool) {
      m, _ = r.m.Get(3)
      return m.Get(1)
   }()
   hash_cash, _ := hash_cash_challenge.Get(1)
   login_context, _ := r.m.GetBytes(5)
   prefix, _ := hash_cash.GetBytes(1)
   length, _ := hash_cash.GetVarint(2)
   if length != 10 {
      return nil, 0, errors.New("invalid hashCash length")
   }
   seed := func() []byte {
      b := sha1.Sum(login_context)
      return b[len(b)-8:]
   }()
   start := time.Now()
   suffix := append(seed, 0,0,0,0,0,0,0,0)
   i := 0
   for {
      input := append(prefix, suffix...)
      digest := sha1.Sum(input)
      if check_ten_trailing_bits(digest[:]) {
         return suffix, i, time.Now().Sub(start), nil
      }
      increment_ctr(suffix, len(suffix)-1)
      increment_ctr(suffix, 7)
      i++
   }
}

func increment_ctr(ctr []byte, index int) {
   ctr[index]++
   if ctr[index] == 0 {
      if index >= 1 {
         increment_ctr(ctr, index-1)
      }
   }
}

func check_ten_trailing_bits(trailing_data []byte) bool {
   length := len(trailing_data)
   if trailing_data[length-1] >= 1 {
      return false
   }
   return count_trailing_zero(trailing_data[length-2]) >= 2
}

type login_response struct {
   m protobuf.Message
}

type login_request struct {
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
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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

func count_trailing_zero(x byte) byte {
   if x == 0 {
      return 32
   }
   var count byte
   for x & 1 == 0 {
      x >>= 1
      count++
   }
   return count
}
