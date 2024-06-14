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
  "included": [
    {
      "attributes": {
        "airDate": "2011-04-01T00:01:00Z",
        "alternateId": "b3b1410a-0c85-457b-bcc7-e13299bea2a8",
        "description": "A soldier is transported into another man’s body aboard a commuter train, where he must foil a terrorist bombing plot within eight minutes.",
        "firstAvailableDate": "2024-04-01T07:01:00Z",
        "isFamilyContent": true,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "A soldier is transported into another man's body on a commuter train, where he must foil a terrorist plot within eight minutes. Learning more about the experimental military technology being deployed, he races against time to save the train passengers.",
        "materialType": "MOVIE",
        "name": "Source Code",
        "originalName": "Source Code",
        "secondaryTitle": "Source Code",
        "videoType": "MOVIE"
      },
      "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8",
      "relationships": {
        "contentPackages": {
          "data": [
            {
              "id": "VodEntertainment",
              "type": "package"
            },
            {
              "id": "Hbo",
              "type": "package"
            }
          ]
        },
        "creditGroups": {
          "data": [
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-starring",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-director",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-writers",
              "type": "creditGroup"
            },
            {
              "id": "b3b1410a-0c85-457b-bcc7-e13299bea2a8-credit-group-producers",
              "type": "creditGroup"
            }
          ]
        },
        "edit": {
          "data": {
            "id": "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
            "type": "edit"
          }
        },
        "images": {
          "data": [
            {
              "id": "i:b",
              "type": "image"
            },
            {
              "id": "i:c",
              "type": "image"
            },
            {
              "id": "i:d",
              "type": "image"
            }
          ]
        },
        "primaryChannel": {
          "data": {
            "id": "c0d1f27a-e2f8-4b3c-bf3c-ed0c4e258093",
            "type": "channel"
          }
        },
        "ratingDescriptors": {
          "data": [
            {
              "id": "4062c194-0c80-4ceb-9de5-d7f7e6e302b8",
              "type": "contentDescriptor"
            },
            {
              "id": "582e60b9-b1f7-4398-9182-a8758856209d",
              "type": "contentDescriptor"
            },
            {
              "id": "f8e36188-362f-4a7b-9934-c7d77cb86e64",
              "type": "contentDescriptor"
            }
          ]
        },
        "ratings": {
          "data": [
            {
              "id": "4d1bf7aa-ddc4-4641-b568-11f59abb4787",
              "type": "contentRating"
            }
          ]
        },
        "show": {
          "data": {
            "id": "127b00c5-0131-4bac-b2d1-40762deefe09",
            "type": "show"
          }
        },
        "txGenres": {
          "data": [
            {
              "id": "0a34f929-1e10-4e53-ba8d-d0e2dbedfdfd",
              "type": "taxonomyNode"
            },
            {
              "id": "9df43ec7-cc1a-4d8f-b7f1-d024810aed66",
              "type": "taxonomyNode"
            },
            {
              "id": "b676b596-7a71-44fb-8222-ecd28becc276",
              "type": "taxonomyNode"
            },
            {
              "id": "cdc65854-4013-4c85-b2b5-c48290ad3b5c",
              "type": "taxonomyNode"
            }
          ]
        }
      },
      "type": "video"
    }
  ],
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
