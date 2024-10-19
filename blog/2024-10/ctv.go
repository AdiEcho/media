package main

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "www.ctv.ca"
   req.URL.Path = "/space-graphql/apq/graphql"
   req.URL.Scheme = "https"
   data := fmt.Sprintf(`
   {
      "operationName": "axisContent",
      "variables": {
         "id": "contentid/axis-content-2968346",
         "subscriptions": [],
         "maturity": "ADULT",
         "language": "ENGLISH",
         "authenticationState": "UNAUTH",
         "playbackLanguage": "ENGLISH"
      },
      "query": %q
   }
   `, data)
   req.Body = io.NopCloser(strings.NewReader(data))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

const data = `
  query axisContent(
    $id: ID!
    $subscriptions: [Subscription]!
    $maturity: Maturity!
    $language: Language!
    $authenticationState: AuthenticationState!
    $playbackLanguage: PlaybackLanguage!
  )
  @uaContext(
    subscriptions: $subscriptions
    maturity: $maturity
    language: $language
    authenticationState: $authenticationState
    playbackLanguage: $playbackLanguage
  ) {
    axisContent(id: $id) {
      axisId
      axisPlaybackLanguages {
        ...AxisPlaybackData
      }
    }
  }

  fragment AxisPlaybackData on AxisPlayback {
    destinationCode
    language
    duration
    playbackIndicators
    partOfMultiLanguagePlayback
  }
`
