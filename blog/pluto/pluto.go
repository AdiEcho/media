package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

type on_demand struct {
   ID string
   Slug string
}

func new_video(slug, forward string) (*on_demand, error) {
   req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName": {"web"},
      "appVersion": {"9"},
      "clientID": {"9"},
      "clientModelNumber": {"9"},
      "episodeSlugs": {slug},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      VOD []on_demand
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   video := s.VOD[0]
   if video.Slug != slug {
      return nil, errors.New(video.Slug)
   }
   return &video, nil
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
   var clips []episode_clip
   err = json.NewDecoder(res.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return &clips[0], nil
}
