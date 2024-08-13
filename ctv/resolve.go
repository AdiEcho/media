package ctv

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type Path string

func (p Path) String() string {
   return string(p)
}

// https://www.ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
// www.ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
// ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
// /shows/friends/the-one-with-the-bullies-s2e21
func (p *Path) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "ctv.ca")
   *p = Path(s)
   return nil
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

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

func (r *ResolvePath) id() string {
   if v := r.FirstPlayableContent; v != nil {
      return v.Id
   }
   return r.Id
}

type ResolvePath struct {
   Id                   string
   FirstPlayableContent *struct {
      Id string
   }
}

func (p Path) Resolve() (*ResolvePath, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         OperationName string `json:"operationName"`
         Query         string `json:"query"`
         Variables     struct {
            Path Path `json:"path"`
         } `json:"variables"`
      }
      s.OperationName = "resolvePath"
      s.Variables.Path = p
      s.Query = graphql_compact(query_resolve)
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   body, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var data struct {
      Data struct {
         ResolvedPath *struct {
            LastSegment struct {
               Content ResolvePath
            }
         }
      }
   }
   err = json.Unmarshal(body, &data)
   if err != nil {
      return nil, err
   }
   if v := data.Data.ResolvedPath; v != nil {
      return &v.LastSegment.Content, nil
   }
   return nil, errors.New(string(body))
}
