package joyn

import (
   "net/http"
   "net/url"
   "os"
   "io"
   "bytes"
   "encoding/json"
)

const page_movie = `
query PageMovieDetailStatic($path: String!) {
   page(path: $path) {
      ... on MoviePage {
         movie {
            ... on Movie {
               ... on Movie {
                  video {
                     id
                  }
               }
            }
         }
      }
   }
}
`

func movie_detail() {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.Query = page_movie
      s.Variables.Path = "/filme/barry-seal-only-in-america"
      return json.Marshal(s)
   }()
   if err != nil {
      panic(err)
   }
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.joyn.de"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Joyn-Platform"] = []string{"web"}
   // x-api-key is hard coded in JavaScript
   req.Header["X-Api-Key"] = []string{"4f0fd9f18abbe3cf0e87fdb556bc39c8"}
   req.Method = "POST"
   req.Body = io.NopCloser(bytes.NewReader(body))
   req.Header.Set("content-type", "application/json")
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
