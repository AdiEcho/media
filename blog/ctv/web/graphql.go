package main

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   // you need this for the first request, then can omit
   req.Header["Graphql-Client-Platform"] = []string{"entpay_web"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.ctv.ca"
   req.URL.Path = "/space-graphql/apq/graphql"
   req.Method = "POST"
   body := map[string]any{
      "operationName": "resolvePath",
      "variables": map[string]any{
         "page":0,
         "path":"/movies/the-girl-with-the-dragon-tattoo-2011",
         "subscriptions":[]string{
            "CTV","CTV_DRAMA","CTV_COMEDY","CTV_LIFE","CTV_SCIFI","CTV_THROWBACK","CTV_MOVIES","CTV_MTV","CTV_MUCH","DISCOVERY","DISCOVERY_SCIENCE","DISCOVERY_VELOCITY","INVESTIGATION_DISCOVERY","ANIMAL_PLANET","E_NOW",
         },
         "maturity":"ADULT",
         "language":"ENGLISH",
         "authenticationState":"UNAUTH",
         "playbackLanguage":"ENGLISH",
      },
      "query": hello,
   }
   text, err := json.Marshal(body)
   if err != nil {
      panic(err)
   }
   req.Body = io.NopCloser(bytes.NewReader(text))
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const hello = `
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
