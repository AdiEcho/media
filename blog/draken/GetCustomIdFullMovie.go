package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "fmt"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODYyMzI3LCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.sYyyMBA7gf0q7A9na8E-vkgJntedFYn2pk_LX2WYBgdQgLgNs7xrtUgR2ZoZlMhgN6D5rQj2U6WDzvDUHZCqEQ"}
   req.Header["Connection"] = []string{"keep-alive"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Host"] = []string{"client-api.magine.com"}
   req.Header["Magine-Accesstoken"] = []string{"22cc71a2-8b77-4819-95b0-8c90f4cf5663"}
   req.Header["Origin"] = []string{"https://drakenfilm.se"}
   req.Header["Referer"] = []string{"https://drakenfilm.se/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"cross-site"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "client-api.magine.com"
   req.URL.Path = "/api/apiql/v2"
   req.URL.Scheme = "https"
   body := fmt.Sprintf(`
   {
      "query": %q,
      "variables": {
         "customId": "the-card-counter",
         "zcustomId": "michael-clayton"
      }
   }
   `, viewable)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const viewable = `
fragment OffersFields on OfferInterfaceType {
  __typename
  ... on DefaultType {
    id
    title
    usps
    image
  }
  ... on RentType {
    id
    title
    usps
    image
    purchaseAvailableUntil
    decorationText
    buttonText
    priceInCents
    currency
  }
  ... on PassType {
    id
    title
    usps
    image
    purchaseAvailableUntil
    decorationText
    buttonText
    priceInCents
    currency
  }
  ... on SubscribeType {
    id
    title
    usps
    image
    trialPeriod {
      length
      unit
    }
    purchaseAvailableUntil
    decorationText
    buttonText
    priceInCents
    currency
  }
}

fragment FullViewableFields on Movie {
  id
  customId
  magineId
  custom
  title
  description
  image: image(type: "poster")
  directors
  genres
  cast
  sixteen: image(type: "sixteen-nine")
  inMyList
  trailer
  productionYear
  duration
  countriesOfOrigin
  defaultPlayable {
    id
    kind
  }
  offers {
    ...OffersFields
  }
  playables {
    ... on VodPlayable {
      id
      kind
      duration
      watchOffset
    }
  }
}

query GetCustomIdFullMovie($customId: ID!) {
  viewer {
    id
    isAuthenticated
    viewableCustomId(customId: $customId) {
      ...FullViewableFields
    }
  }
}
`
