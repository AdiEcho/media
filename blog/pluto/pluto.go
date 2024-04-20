package pluto

import (
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "strings"
)

var forwards = map[string]string{"Canada": "99.224.0.0"}

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
      "drmCapabilities": {"widevine:L3"},
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
   if video.ID != slug {
      if video.Slug != slug {
         return nil, fmt.Errorf("%+v", video)
      }
   }
   return &video, nil
}

var bases = []url.URL{
   // these return `403 OK` with compressed content
   {Scheme: "http", Host: "siloh-fs.plutotv.net"},
   {Scheme: "http", Host: "siloh-ns1.plutotv.net"},
   {Scheme: "https", Host: "siloh-fs.plutotv.net"},
   {Scheme: "https", Host: "siloh-ns1.plutotv.net"},
   // returns `200 OK` with plain content
   {Scheme: "http", Host: "silo-hybrik.pluto.tv.s3.amazonaws.com"},
}

func (s source) parse(base url.URL) (*url.URL, error) {
   ref, err := url.Parse(s.File)
   if err != nil {
      return nil, err
   }
   ref.Scheme = base.Scheme
   ref.Host = base.Host
   return ref, nil
}

type source struct {
   File string
   Type string
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

////////

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
