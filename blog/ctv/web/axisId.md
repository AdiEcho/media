# axisId

~~~
"\n  fragment AuthConstraintsData on AuthConstraint {\n    authRequired\n    packageName\n    endDate\n    language\n    startDate\n    subscriptionName\n    __typename\n  }\n"
"\n  query resolvePath(\n    $path: String!\n    $subscriptions: [Subscription]!\n    $maturity: Maturity!\n    $language: Language!\n    $authenticationState: AuthenticationState!\n    $playbackLanguage: PlaybackLanguage!\n    $page: Int = 0\n  )\n  @uaContext(\n    subscriptions: $subscriptions\n    maturity: $maturity\n    language: $language\n    authenticationState: $authenticationState\n    playbackLanguage: $playbackLanguage\n  ) {\n    resolvedPath(path: $path) {\n      redirected\n      path\n      segments {\n        position\n        content {\n          title\n          id\n          path\n        }\n      }\n      lastSegment {\n        position\n        content {\n          id\n          title\n          path\n          __typename\n          ... on AceWebContent {\n            path\n            ... on Rotator {\n              adUnit {\n                ...AceAdUnitData\n              }\n            }\n\n            ... on Article {\n              seoFields {\n                seoDescription\n                seoTitle\n                canonicalUrl\n              }\n              ogFields {\n                ...OGFields\n              }\n            }\n          }\n          ... on AxisObject {\n            __typename\n            description\n            axisId\n            ... on AxisContent {\n              keywords\n              adUnit {\n                ...AxisAdUnitData\n              }\n              seasonNumber\n              episodeNumber\n              contentType\n              seoFields {\n                ...SEOFields\n              }\n              ogFields {\n                ...OGFields\n              }\n            }\n            ... on AxisCollection {\n              adUnit {\n                ...AxisAdUnitData\n              }\n              seoFields {\n                ...SEOFields\n              }\n            }\n            ... on AxisMedia {\n              keywords\n              adUnit {\n                ...AxisAdUnitData\n              }\n              firstPlayableContent {\n                id\n                axisId\n                authConstraints {\n                  ...AuthConstraintsData\n                }\n                axisPlaybackLanguages {\n                  ...AxisPlaybackData\n                }\n                badges {\n                  title\n                  label\n                }\n              }\n              seoFields {\n                ...SEOFields\n              }\n              ogFields {\n                ...OGFields\n              }\n            }\n          }\n          ... on Section {\n            containerType\n            adUnit {\n              ...AceAdUnitData\n            }\n            gridConfig {\n              ...GridConfigData\n            }\n            secondNavigation {\n              title\n              renderTitleAs\n              titleImage {\n                __typename\n                id\n                url\n              }\n              links {\n                ...LinkData\n              }\n            }\n            firstPageLayout {\n              ...FirstPageLayoutData\n            }\n            seoFields {\n              ...SEOFields\n            }\n            ogFields {\n              ...OGFields\n            }\n          }\n          ... on Site {\n            adUnit {\n              ...AceAdUnitData\n            }\n            firstPageLayout {\n              ...FirstPageLayoutData\n            }\n            siteConfig {\n              ...SiteConfig\n            }\n            seoFields {\n              ...SEOFields\n            }\n            ogFields {\n              ...OGFields\n            }\n          }\n        }\n      }\n      searchResults {\n        ... on Medias {\n          page(page: $page) {\n            totalItemCount\n            totalPageCount\n            hasNextPage\n            items {\n              displayGenres @client\n              displayWatchList @client\n              id\n              title\n              summary\n              agvotCode\n              qfrCode\n              axisId\n              path\n              metadataUpgrade {\n                ...AxisMetaDataUpgradeData\n              }\n              posterImages: images(formats: POSTER) {\n                url\n              }\n              squareImages: images(formats: SQUARE) {\n                url\n              }\n              thumbnailImages: images(formats: THUMBNAIL) {\n                url\n              }\n              firstPlayableContent {\n                ...FirstPlayableContentData\n              }\n              badges {\n                title\n                label\n              }\n              genres {\n                name\n              }\n              firstAirYear\n              originatingNetworkLogoId\n              heroBrandLogoId\n              seasons {\n                id\n              }\n            }\n          }\n        }\n        ... on Articles {\n          page(page: $page) {\n            totalItemCount\n            totalPageCount\n            hasNextPage\n            items {\n              id\n              title\n              path\n            }\n          }\n        }\n      }\n    }\n  }\n  \n  \n  \n  \n  \n  \n  \n  \n  \n  \n  \n  \n"
~~~

## movies

this request:

~~~
POST https://www.ctv.ca/space-graphql/apq/graphql HTTP/2.0
graphql-client-platform: entpay_web

{
 "operationName": "resolvePath",
 "query": "query resolvePath($path: String!) { resolvedPath(path: $path) { lastSegment { content { ... on AxisObject { axisId ... on AxisMedia { firstPlayableContent { axisId } } } } } } }",
 "variables": {
  "path": "/movies/the-girl-with-the-dragon-tattoo-2011"
 }
}
~~~

has this response body:

~~~json
{
  "data": {
    "resolvedPath": {
      "lastSegment": {
        "content": {
          "axisId": 43239,
          "firstPlayableContent": {
            "axisId": 1417780
          }
        }
      }
    }
  }
}
~~~

these work:

- <https://capi.9c9media.com/destinations/ctvmovies_hub/platforms/desktop/contents/1417780/contentPackages>
- <https://capi.9c9media.com/destinations/ctvmovies_hub/platforms/desktop/contents/1417780?$include=[ContentPackages]>

these fail:

- <https://capi.9c9media.com/destinations/ctvmovies_hub/platforms/desktop/contents/43239/contentPackages>
- <https://capi.9c9media.com/destinations/ctvmovies_hub/platforms/desktop/contents/43239?$include=[ContentPackages]>

## shows

this request:

~~~
POST https://www.ctv.ca/space-graphql/apq/graphql HTTP/2.0
graphql-client-platform: entpay_web

{
 "operationName": "resolvePath",
 "query": "query resolvePath($path: String!) { resolvedPath(path: $path) { lastSegment { content { ... on AxisObject { axisId ... on AxisMedia { firstPlayableContent { axisId } } } } } } }",
 "variables": {
  "path": "/shows/friends/the-one-with-the-bullies-s2e21"
 }
}
~~~

has this response body:

~~~json
{
  "data": {
    "resolvedPath": {
      "lastSegment": {
        "content": {
          "axisId": 1730820
        }
      }
    }
  }
}
~~~

these work:

- <https://capi.9c9media.com/destinations/ctvcomedy_hub/platforms/desktop/contents/1730821/contentPackages>
- <https://capi.9c9media.com/destinations/ctvcomedy_hub/platforms/desktop/contents/1730821?$include=[ContentPackages]>
