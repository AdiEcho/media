package ctv

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               ... on AxisMedia {
                  firstPlayableContent {
                     axisId
                  }
               }
            }
         }
      }
   }
}
`

type playable_content struct {
   AxisId int64
}

func first_playable_content(path string) (*playable_content, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         OperationName string `json:"operationName"`
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.OperationName = "resolvePath"
      s.Query = query_resolve
      s.Variables.Path = path
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Data struct {
         ResolvedPath struct {
            LastSegment struct {
               Content struct {
                  FirstPlayableContent playable_content
               }
            }
         }
      }
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Data.ResolvedPath.LastSegment.Content.FirstPlayableContent, nil
}

func (p playable_content) items() (*content_packages, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, p.AxisId, 10)
      b = append(b, "/contentPackages"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var items content_packages
   items.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &items, nil
}

type content_package struct {
   ID int64
}

type content_packages struct {
   data []byte
   v struct {
      Items []content_package
   }
}

func (c content_packages) item() (*content_package, error) {
   err := json.Unmarshal(c.data, &c.v)
   if err != nil {
      return nil, err
   }
   return &c.v.Items[0], nil
}
