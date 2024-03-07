package android

import (
   "154.pages.dev/protobuf"
   "bytes"
   "encoding/hex"
   "errors"
   "io"
   "net/http"
)

func (f file_format) file_id() (string, bool) {
   if v, ok := f.m.GetBytes(1); ok {
      return hex.EncodeToString(v), true
   }
   return "", false
}

func (f file_format) OGG_VORBIS_320() bool {
   if v, ok := f.format(); ok {
      if v == 2 {
         return true
      }
   }
   return false
}

func (f file_format) format() (uint64, bool) {
   if v, ok := f.m.GetVarint(2); ok {
      return uint64(v), true
   }
   return 0, false
}

type file_format struct {
   m protobuf.Message
}

func (m metadata) file() []file_format {
   var vs []file_format
   for _, field := range m.m {
      if v, ok := field.Get(2); ok {
         if v, ok := v.Get(3); ok {
            if v, ok := v.Get(3); ok {
               if v, ok := v.Get(2); ok {
                  for _, field := range v {
                     if v, ok := field.Get(12); ok {
                        vs = append(vs, file_format{v})
                     }
                  }
               }
            }
         }
      }
   }
   return vs
}

// github.com/librespot-org/librespot/blob/dev/protocol/proto/media_format.proto
// github.com/librespot-org/librespot/blob/dev/protocol/proto/metadata.proto
// github.com/librespot-org/librespot/blob/dev/protocol/proto/media_manifest.proto
type metadata struct {
   m protobuf.Message
}

func (o LoginOk) metadata(canonical_uri string) (*metadata, error) {
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
   var meta metadata
   if err := meta.m.Consume(data); err != nil {
      return nil, err
   }
   return &meta, nil
}
