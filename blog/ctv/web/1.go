package ctv

import (
   "bytes"
   "encoding/json"
   "net/http"
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
      s.Query = query_resolve
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

type resolve_path struct {
   Data struct {
      ResolvedPath struct {
         LastSegment struct {
            Content struct {
               AxisId int64
               FirstPlayableContent struct {
                  AxisId int64
               }
            }
         }
      }
   }
}
