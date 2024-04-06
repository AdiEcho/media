package plex

import (
   "net/http"
   "net/url"
   "strings"
)

func (a anonymous) metadata(address string) (*http.Response, error) {
   req, err := http.NewRequest("GET", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "X-Plex-Token": {a.AuthToken},
   }
   req.URL.Path, err = func() (string, error) {
      u, err := url.Parse(address)
      if err != nil {
         return "", err
      }
      u.Path = strings.Replace(u.Path, "/movie/", "/movie:", 1)
      return "/library/metadata" + u.Path, nil
   }()
   if err != nil {
      return nil, err
   }
   return http.DefaultClient.Do(req)
}
