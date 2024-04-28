# axisId

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
