# destinations

## shows

this request:

~~~
POST https://www.ctv.ca/space-graphql/apq/graphql HTTP/2.0
graphql-client-platform: entpay_web

{
 "operationName": "resolvePath",
 "query": "query resolvePath($path: String!) { resolvedPath(path: $path) { lastSegment { content { ... on AxisObject { axisId ... on AxisMedia { firstPlayableContent { authConstraints { ... on AuthConstraint { packageName } } axisId axisPlaybackLanguages { ... on AxisPlayback { destinationCode } } } } } } } } }",
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

## movies

this request:

~~~
POST https://www.ctv.ca/space-graphql/apq/graphql HTTP/2.0
graphql-client-platform: entpay_web

{
 "operationName": "resolvePath",
 "query": "query resolvePath($path: String!) { resolvedPath(path: $path) { lastSegment { content { ... on AxisObject { axisId ... on AxisMedia { firstPlayableContent { authConstraints { ... on AuthConstraint { packageName } } axisId axisPlaybackLanguages { ... on AxisPlayback { destinationCode } } } } } } } } }",
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
            "authConstraints": [
              {
                "packageName": "ctvmovies_hub"
              }
            ],
            "axisId": 1417780,
            "axisPlaybackLanguages": [
              {
                "destinationCode": "ctvmovies_hub"
              }
            ]
          }
        }
      }
    }
  }
}
~~~
