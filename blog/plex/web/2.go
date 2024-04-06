package plex

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a anonymous) metadata(address string) (*metadata, error) {
   match, err := url.Parse(address)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "GET", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "X-Plex-Token": {a.AuthToken},
      "url": {match.Path},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      MediaContainer struct {
         Metadata []metadata
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

type metadata struct {
   Media []struct {
      Part []struct {
         Key string
         License string
      }
   }
}

//req.URL.Host = "vod.provider.plex.tv"
//req.URL.Path = "/library/parts/64cc0e5a7a36935c7ba4eb96-dash/license"
//req.URL.Scheme = "https"
//val["X-Plex-DRM"] = []string{"widevine"}
//val["X-Plex-Token"] = []string{"fc1WPqnLdmq3J4Axt5pn"}
func (metadata) RequestUrl() (string, bool) {
   return "", false
}

func (metadata) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (metadata) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (metadata) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}
