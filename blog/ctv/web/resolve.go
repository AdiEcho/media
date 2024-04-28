package ctv

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               axisId
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
