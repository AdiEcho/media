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
  fragment AuthConstraintsData on AuthConstraint {
    authRequired
    packageName
    endDate
    language
    startDate
    subscriptionName
    __typename
  }

  fragment AxisAdUnitData on AxisAdUnit {
    adultAudience
    heroBrand
    pageType
    product
    revShare
    title
    analyticsTitle
    keyValue {
      webformType
      adTarget
      contentType
      mediaType
      pageTitle
      revShare
      subType
    }
  }

  fragment AxisPlaybackData on AxisPlayback {
    destinationCode
    language
    duration
    playbackIndicators
    partOfMultiLanguagePlayback
  }

  fragment LinkData on Link {
 buttonStyle
    urlParameters
    renderAs
    linkType
    linkLabel
    longLinkLabel
    linkTarget
    userMgmtLinkType
    url
    id
    showLinkLabel
    internalContent {
      title
      __typename
      ... on AxisContent {
        axisId
        authConstraints {
          ...AuthConstraintsData
        }
        agvotCode
      }
      ... on AceWebContent {
        path
        pathSegment
        __typename
      }
      ... on Section {
        containerType
        path
      }
      ... on AxisObject {
        axisId
        title
      }
      ... on TabItem {
        # id
        sectionPath
      }
    }
    hoverImage {
      title
      imageType
      url
    }
    image {
      id
      width
      height
      title
      url
      altText
    }
    bannerImages {
      breakPoint
 image {
        id
        title
        url
        altText
      }
    }
    __typename
  }
  

  fragment RotatorConfigData on RotatorConfig {
    displayTitle
    displayTotalItemCount
    displayDots
    style
    imageFormat
    lightbox
    carousel
    titleLinkMode
    maxItems
    disableBadges
    customTitleLink {
      ...LinkData
    }
    hideMediaTitle
  }
  

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
      id
      path
      title
      duration
      agvotCode
      description
      episodeNumber
      seasonNumber
      pathSegment
      genres {
        name
      }
      axisMedia {
        heroBrandLogoId
        id
        title
      }
      adUnit {
        ...AxisAdUnitData
      }
      authConstraints {
        ...AuthConstraintsData
      }
      axisPlaybackLanguages {
        ...AxisPlaybackData
      }
      originalSpokenLanguage
      ogFields {
        ogDescription
        ogImages {
          url
        }
        ogTitle
      }
      playbackMetadata {
        indicator
        languages {
          languageCode
          languageDisplayName
        }
      }
      seoFields {
        seoDescription
        seoTitle
        seoKeywords
        canonicalUrl
      }
      badges {
        title
        label
      }
      posterImages: images(formats: POSTER) {
        url
      }
      broadcastDate
      expiresOn
      startsOn
      keywords
      videoPageLayout {
        __typename
        ... on Rotator {
          id
          config {
            ...RotatorConfigData
          }
        }
      }
    }
  }
`
