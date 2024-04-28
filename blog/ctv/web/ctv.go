package ctv

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
   "strings"
)

/*
for shows this is:
"id": "contentid/axis-content-1730820"

for movies this is:
"firstPlayableContent": {
"id": "contentid/axis-content-1417780"
*/
const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               id
               ... on AxisMedia {
                  firstPlayableContent {
                     id
                  }
               }
            }
         }
      }
   }
}
`

type resolve_path struct {
   data []byte
   v struct {
      Data struct {
         ResolvedPath struct {
            LastSegment last_segment
         }
      }
   }
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

func (r *resolve_path) New(path string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         OperationName string `json:"operationName"`
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.OperationName = "resolvePath"
      s.Variables.Path = path
      s.Query = graphql_compact(query_resolve)
      return json.MarshalIndent(s, "", " ")
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   r.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r *resolve_path) unmarshal() error {
   return json.Unmarshal(r.data, &r.v)
}
type content_packages struct {
   data []byte
   v struct {
      Items []struct {
         ID int64
      }
   }
}

func (c *content_packages) unmarshal() error {
   return json.Unmarshal(c.data, &c.v)
}

type poster struct{}

func (poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (poster) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) RequestUrl() (string, bool) {
   return "https://license.9c9media.ca/widevine", true
}

type last_segment struct {
   Content struct {
      AxisId int64
      FirstPlayableContent struct {
         AxisId int64
      }
   }
}

// wikipedia.org/wiki/Geo-blocking
func (s last_segment) manifest(c content_packages) string {
   b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
   b = append(b, "/platforms/desktop/playback/contents/"...)
   b = strconv.AppendInt(b, s.Content.FirstPlayableContent.AxisId, 10)
   b = append(b, "/contentPackages/"...)
   b = strconv.AppendInt(b, c.v.Items[0].ID, 10)
   b = append(b, "/manifest.mpd"...)
   return string(b)
}

func (s last_segment) packages() (*content_packages, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, s.Content.FirstPlayableContent.AxisId, 10)
      b = append(b, "/contentPackages"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var packages content_packages
   packages.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &packages, nil
}
