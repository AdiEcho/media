package pluto

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

// `https://siloh.pluto.tv` returns `403 Forbidden` with no content.
// `http://silo-hybrik.pluto.tv.s3.amazonaws.com` returns `200 OK` with plain
// content. `https://siloh-fs.plutotv.net` returns `403 OK` with compressed
// content.
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

func (o on_demand) clip() (*episode_clip, error) {
   req, err := http.NewRequest("GET", "http://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      b.WriteString(o.ID)
      b.WriteString("/clips.json")
      return b.String()
   }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var clips []episode_clip
   err = json.NewDecoder(res.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return &clips[0], nil
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

func (e episode_clip) hls() (*source, bool) {
   for _, s := range e.Sources {
      if s.Type == "HLS" {
         return &s, true
      }
   }
   return nil, false
}
