package ctv

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
)

type content_packages struct {
   Items []struct {
      ID int
   }
}

func (r resolve_path) content() (*content_packages, error) {
   address := func() string {
      b := []byte("https://capi.9c9media.com/destinations/ctvmovies_hub")
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, r.Data.ResolvedPath.LastSegment.Content.FirstPlayableContent.AxisId, 10)
      b = append(b, "/contentPackages"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   content := new(content_packages)
   err = json.NewDecoder(res.Body).Decode(content)
   if err != nil {
      return nil, err
   }
   return content, nil
}

type resolve_path struct {
  Data struct {
    ResolvedPath struct {
      LastSegment struct {
        Content struct {
          FirstPlayableContent struct {
            AxisId int64
          }
        }
      }
    }
  }
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
      s.Query = query_resolve
      s.Variables.Path = path
      return json.Marshal(s)
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
   return json.NewDecoder(res.Body).Decode(r)
}

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
