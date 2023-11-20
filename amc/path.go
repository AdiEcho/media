package amc

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (p Path) nid() (string, error) {
   _, nid, found := strings.Cut(p.s, "--")
   if !found {
      return "", errors.New("nid")
   }
   return nid, nil
}

type Path struct {
   s string
}

func (p Path) String() string {
   return p.s
}

func (p *Path) Set(s string) error {
   if _, after, found := strings.Cut(s, "://"); found {
      s = after // remove scheme
   }
   if i := strings.IndexByte(s, '/'); i >= 1 {
      s = s[i:] // remove host
   }
   p.s = s
   return nil
}

func (a Auth_ID) Content(p Path) (*Content, error) {
   req, err := http.NewRequest("GET", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   req.URL.Path, err = func() (string, error) {
      u, err := url.Parse(p.s)
      if err != nil {
         return "", err
      }
      var b strings.Builder
      b.WriteString("/content-compiler-cr/api/v1/content/amcn/amcplus/path")
      if strings.HasPrefix(u.Path, "/movies/") {
         b.WriteString("/watch")
      }
      b.WriteString(u.Path)
      return b.String(), nil
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}
