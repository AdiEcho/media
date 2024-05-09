package main

import (
   "net/http"
   "net/url"
   "os"
   "fmt"
   "io"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.joyn.de"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   req.Method = "POST"
   req.Header["Joyn-Platform"] = []string{"web"}
   req.Header["X-Api-Key"] = []string{"4f0fd9f18abbe3cf0e87fdb556bc39c8"}
   body := fmt.Sprintf(`
   {
      "variables": {
         "path": "/serien/one-tree-hill/1-2-quaelende-angst"
      },
      "query": %q
   }
   `, episode_detail)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const episode_detail = `
query EpisodeDetailPageStatic($path: String!) {
  page(path: $path) {
      __typename
      path
      ... on EpisodePage {
          episode {
              ... on Episode {
                 number
                 title
                 season {
                     ... on Season {
                       number
                     }
                 }
                 video {
                     id
                 }
                 series {
                     title
                     
                     id
                     subtype
                     path
                     productionYear
                     productionCompanies
                     productionCountries
                     copyrights
                     path
                     numberOfSeasons
                 }
                 __typename
                 id
                 path
                 airdate
                 startsAt
                 ageRating {
                     minAge
                     descriptorsText
                 }
                 licenseTypes
                 description
                 brands {
                     path
                 }
                 markings
                 genres {
                     name
                 }
                 videoDescriptors {
                     name
                 }
                 productPlacement
               }
          }
      }
  }
}
`
