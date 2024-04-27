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
  query resolvePath(
    $path: String!
    $subscriptions: [Subscription]!
    $maturity: Maturity!
    $language: Language!
    $authenticationState: AuthenticationState!
    $playbackLanguage: PlaybackLanguage!
    $page: Int = 0
  )
  @uaContext(
    subscriptions: $subscriptions
    maturity: $maturity
    language: $language
    authenticationState: $authenticationState
    playbackLanguage: $playbackLanguage
  ) {
    resolvedPath(path: $path) {
      redirected
      path
      segments {
        position
        content {
          title
          id
          path
        }
      }
      lastSegment {
        position
        content {
          id
          title
          path
          __typename
          ... on AceWebContent {
            path
          }
          ... on AxisObject {
            __typename
            description
            axisId
            ... on AxisContent {
              keywords
              seasonNumber
              episodeNumber
              contentType
            }
            ... on AxisMedia {
              keywords
              firstPlayableContent {
                id
                axisId
                badges {
                  title
                  label
                }
              }
            }
          }
          ... on Section {
            containerType
            secondNavigation {
              title
              renderTitleAs
              titleImage {
                __typename
                id
                url
              }
            }
          }
        }
      }
      searchResults {
        ... on Medias {
          page(page: $page) {
            totalItemCount
            totalPageCount
            hasNextPage
            items {
              id
              title
              summary
              agvotCode
              qfrCode
              axisId
              path
              posterImages: images(formats: POSTER) {
                url
              }
              squareImages: images(formats: SQUARE) {
                url
              }
              thumbnailImages: images(formats: THUMBNAIL) {
                url
              }
              badges {
                title
                label
              }
              genres {
                name
              }
              firstAirYear
              originatingNetworkLogoId
              heroBrandLogoId
              seasons {
                id
              }
            }
          }
        }
        ... on Articles {
          page(page: $page) {
            totalItemCount
            totalPageCount
            hasNextPage
            items {
              id
              title
              path
            }
          }
        }
      }
    }
  }
`
