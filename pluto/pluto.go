package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

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

type source struct {
   File string
   Type string
}

var bases = []string{
   // these return `403 OK` with compressed content
   "http://siloh-fs.plutotv.net",
   "http://siloh-ns1.plutotv.net",
   "https://siloh-fs.plutotv.net",
   "https://siloh-ns1.plutotv.net",
   // returns `200 OK` with plain content
   "http://silo-hybrik.pluto.tv.s3.amazonaws.com",
}

func (s source) parse(base string) (*url.URL, error) {
   a, err := url.Parse(base)
   if err != nil {
      return nil, err
   }
   b, err := url.Parse(s.File)
   if err != nil {
      return nil, err
   }
   b.Scheme = a.Scheme
   b.Host = a.Host
   return b, nil
}
