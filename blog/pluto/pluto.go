package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func new_clip(id string) (*episode_clip, error) {
   req, err := http.NewRequest("GET", "http://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      b.WriteString(id)
      b.WriteString("/clips.json")
      return b.String()
   }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var clips []episode_clip
   err = json.NewDecoder(res.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return &clips[0], nil
}

type poster struct{}

func (poster) RequestUrl() (string, bool) {
   return "https://service-concierge.clusters.pluto.tv/v1/wv/alt", true
}

func (poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (poster) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

// `https://siloh.pluto.tv` returns `403 Forbidden` with no content.
// `http://silo-hybrik.pluto.tv.s3.amazonaws.com` returns `200 OK` with plain
// content.
// `https://siloh-fs.plutotv.net` returns `403 OK` with compressed content.
// `https://siloh-ns1.plutotv.net` returns `403 OK` with compressed content.
func (s source) parse() (*url.URL, error) {
   u, err := url.Parse(s.File)
   if err != nil {
      return nil, err
   }
   u.Host = "siloh-fs.plutotv.net"
   return u, nil
}

type source struct {
   File string
   Type string
}

type episode_clip struct {
   Sources []source
}

func (e episode_clip) dash() (*source, bool) {
   for _, s := range e.Sources {
      if s.Type == "DASH" {
         return &s, true
      }
   }
   return nil, false
}
