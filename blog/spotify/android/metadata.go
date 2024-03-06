package android

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

/*
github.com/librespot-org/librespot/blob/dev/protocol/proto/media_format.proto
1 "format": "OGG_VORBIS_96"
2 "format": "OGG_VORBIS_160"
3 "format": "OGG_VORBIS_320"
9 "format": "AAC_24"
11 "format": "MP4_128"
12 "format": "MP4_128_DUAL"
14 "format": "MP4_256"
15 "format": "MP4_256_DUAL"

github.com/librespot-org/librespot/blob/dev/protocol/proto/metadata.proto
github.com/librespot-org/librespot/blob/dev/protocol/proto/media_manifest.proto
0 "format": "OGG_VORBIS_96"
1 "format": "OGG_VORBIS_160"
2 "format": "OGG_VORBIS_320"
8 "format": "AAC_24"
"format": "MP4_128"
"format": "MP4_128_DUAL"
"format": "MP4_256"
"format": "MP4_256_DUAL"
*/
func (o LoginOk) metadata(canonical_uri string) (protobuf.Message, error) {
   token, ok := o.AccessToken()
   if !ok {
      return nil, errors.New("LoginOk.AccessToken")
   }
   var m protobuf.Message
   m.Add(2, func(m *protobuf.Message) {
      m.AddBytes(1, []byte(canonical_uri))
      m.Add(2, func(m *protobuf.Message) {
         m.AddVarint(1, 10)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://guc3-spclient.spotify.com", bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/extended-metadata/v0/extended-metadata"
   req.Header.Set("authorization", "Bearer "+token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var message protobuf.Message
   if err := message.Consume(data); err != nil {
      return nil, err
   }
   return message, nil
}
