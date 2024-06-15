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

movie:

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
        "edit": {
          "data": {
            "id": "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
            "type": "edit"
          }
        },
        "primaryChannel": {
          "data": {
            "id": "c0d1f27a-e2f8-4b3c-bf3c-ed0c4e258093",
            "type": "channel"
          }
        },
        "show": {
          "data": {
            "id": "127b00c5-0131-4bac-b2d1-40762deefe09",
            "type": "show"
          }
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

episode:

~~~json
{
  "data": {
    "attributes": {
      "canonical": true,
      "url": "/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68"
    },
    "id": "83b758280df91c09cd2e3fd2e2a2013b90959854ca2ae772de1caa9942578e26",
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
        "airDate": "2000-01-01T00:01:00Z",
        "alternateId": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f",
        "description": "Tony travels to Italy to jump-start a car-importing \"business\" and recruits a new lieutenant named Furio.",
        "episodeNumber": 4,
        "firstAvailableDate": "2023-02-18T05:00:00Z",
        "isFamilyContent": false,
        "isFavorite": false,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "Tony travels to Italy to jump-start a car-importing 'business' and recruits a new lieutenant named Furio.",
        "materialType": "EPISODE",
        "name": "Commendatori",
        "originalName": "Commendatori",
        "seasonNumber": 2,
        "secondaryTitle": "Commendatori",
        "videoType": "EPISODE",
        "viewingHistory": {
          "completed": false,
          "lastReportedTimestamp": "2024-06-14T00:14:28.083Z",
          "position": 51,
          "viewed": true
        }
      },
      "id": "fbdd33a2-1189-4b9a-8c10-13244fb21b7f",
      "relationships": {
        "show": {
          "data": {
            "id": "818c3d9d-1831-48a6-9583-0364a7f98453",
            "type": "show"
          }
        }
      },
      "type": "video"
    },
    {
      "attributes": {
        "alternateId": "818c3d9d-1831-48a6-9583-0364a7f98453",
        "description": "James Gandolfini stars in this acclaimed series about a mob boss whose professional and private strains land him in therapy.",
        "firstAvailableDate": "2023-02-08T05:00:00Z",
        "isFamilyContent": false,
        "isFavorite": false,
        "isKidsContent": false,
        "kidsContent": false,
        "longDescription": "James Gandolfini stars in this acclaimed series about a mob boss whose professional and private strains land him in therapy.",
        "name": "The Sopranos",
        "numberOfNewEpisodes": 0,
        "originalName": "The Sopranos",
        "premiereDate": "1999-01-21T00:01:00Z",
        "showType": "SERIES"
      },
      "id": "818c3d9d-1831-48a6-9583-0364a7f98453",
      "type": "show"
    },
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
