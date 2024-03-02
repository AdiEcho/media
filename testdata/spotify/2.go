package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

func (r login_response) challenge_solution(
   suffix []byte, iterations int,
) (protobuf.Message, error) {
   var m protobuf.Message
   m.AddFunc(1, func(m *protobuf.Message) {
      m.AddBytes(1, []byte("9a8d2f0ce77a4e248bb71fefcb557637"))
      m.AddBytes(2, []byte("58cebdd226ac462a"))
   })
   login_context, _ := r.m.GetBytes(5)
   m.AddBytes(2, login_context)
   m.AddFunc(3, func(m *protobuf.Message) {
      m.AddFunc(1, func(m *protobuf.Message) {
         m.AddFunc(1, func(m *protobuf.Message) {
            m.AddBytes(1, suffix)
            m.AddFunc(2, func(m *protobuf.Message) {
               m.AddVarint(2, protobuf.Varint(iterations))
            })
         })
      })
   })
   req, err := http.NewRequest(
      "POST", "https://login5.spotify.com/v3/login", bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Cache-Control": {"no-cache, no-store, max-age=0"},
      "Client-Token": {"AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q=="},
      "Content-Type": {"application/x-protobuf"},
      "User-Agent": {"Spotify/8.9.18.512 Android/23 (Android SDK built for x86)"},
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
   m = nil
   if err := m.Consume(data); err != nil {
      return nil, err
   }
   return m, nil
}
