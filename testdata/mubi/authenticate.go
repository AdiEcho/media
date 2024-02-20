package mubi

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (c linkCode) authenticate() (*http.Response, error) {
   body, err := json.Marshal(map[string]string{"auth_token": c.s.Auth_Token})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.mubi.com/v3/authenticate", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "Accept-Language": {"en-US"},
      "Client": {"android"},
      "Client-Accept-Audio-Codecs": {"AAC"},
      "Client-App": {"mubi"},
      "Client-Country": {"US"},
      "Client-Device-Brand": {"unknown"},
      "Client-Device-Identifier": {"437cacfa-7421-410c-a4cd-fbe5d5460345"},
      "Client-Device-Model": {"Android SDK built for x86"},
      "Client-Device-Os": {"6.0"},
      "Client-Version": {"41.2"},
      "Content-Type": {"application/json; charset=UTF-8"},
      "User-Agent": {"okhttp/4.10.0"},
   }
   return http.DefaultClient.Do(req)
}
