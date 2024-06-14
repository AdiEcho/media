# attributes

start here:

https://play.max.com/movie/127b00c5-0131-4bac-b2d1-40762deefe09

then:

https://play.max.com/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579

then:

~~~
GET https://default.any-amer.prd.api.discomax.com/cms/routes/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579?include=default HTTP/2.0
authorization: Bearer eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi0yNTIzYWEyNC1jNzU...
~~~

result:

~~~json
{
  "data": {
    "attributes": {
      "canonical": true,
      "url": "/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579"
    },
    "id": "c65fc42b4caa4d290d40a9c3fcc192a4adf064b9870b677aa8aab2f8510072f1",
    "relationships": {
      "target": {
        "data": {
          "id": "215925083836394013750542481199525881753",
          "type": "page"
        }
      }
    },
    "type": "route"
  },
  "meta": {
    "site": {
      "attributes": {
        "brandId": "beam",
        "mainTerritoryCode": "us",
        "theme": "beam",
        "websiteHostName": ""
      },
      "id": "beam_us",
      "type": "site"
    }
  }
}
~~~
