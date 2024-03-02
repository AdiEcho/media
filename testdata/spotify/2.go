package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
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
   suffix := solve_hash_cash_challenge(login_context, prefix)
   var m protobuf.Message
   m.AddFunc(1, func(m *protobuf.Message) {
      m.AddBytes(1, []byte("9a8d2f0ce77a4e248bb71fefcb557637"))
   })
   m.AddBytes(2, login_context)
   m.AddFunc(3, func(m *protobuf.Message) {
      m.AddFunc(1, func(m *protobuf.Message) {
         m.AddFunc(1, func(m *protobuf.Message) {
            m.AddBytes(1, suffix)
         })
      })
   })
   m.AddFunc(101, func(m *protobuf.Message) {
      m.AddBytes(1, []byte(username))
      m.AddBytes(2, []byte(password))
   })
   req, err := http.NewRequest(
      "POST", "https://login5.spotify.com/v3/login", bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-protobuf"},
      //"Cache-Control": {"no-cache, no-store, max-age=0"},
      //"User-Agent": {"Spotify/8.9.18.512 Android/23 (Android SDK built for x86)"},
      "User-Agent": {"Symfony HttpClient (Curl)"},
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
   m = nil
   if err := m.Consume(data); err != nil {
      return nil, err
   }
   return m, nil
}
