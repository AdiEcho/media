package main

import (
   "net/http"
   "net/url"
   "os"
   "io"
   "bytes"
   "encoding/json"
)

func main() {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.Query = movie_detail
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

const movie_detail = `
query PageMovieDetailStatic($path: String!) {
  page(path: $path) {
      ... on MoviePage {
          path
          tracking {
              pageName
              payload
          }
          id
          movie {
              ... on Movie {
                 id
                 __typename
                 tagline
                 path
                 licenseTypes
                 markings
                 licenseTypes
                 tracking {
                     agofCode
                     externalAssetId
                 }
                 ... on Movie {
                    title
                    ageRating {
                        minAge
                        descriptorsText
                    }
                    brands {
                        path
                    }
                    copyrights
                    description
                    languages {
                        code
                        name
                    }
                    productionYear
                    productionCountries
                    productPlacement
                    video {
                        id
                        duration
                        audioLanguages {
                            name
                        }
                        quality
                    }
                    markings
                  }
               }

          }
      }
  }
}
`
