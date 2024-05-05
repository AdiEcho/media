package draken

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strings"
   "fmt"
)

type full_movie struct {
   Data struct {
      Viewer struct {
         ViewableCustomId struct {
            DefaultPlayable struct {
               ID string
            }
         }
      }
   }
}

func (f *full_movie) New() error {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "client-api.magine.com"
   req.URL.Path = "/api/apiql/v2"
   req.URL.Scheme = "https"
   req.Header["x-forwarded-for"] = []string{"78.64.0.0"}
   req.Header["Magine-Accesstoken"] = []string{"22cc71a2-8b77-4819-95b0-8c90f4cf5663"}
   body := fmt.Sprintf(`
   {
      "query": %q,
      "variables": {
         "customId": "michael-clayton"
      }
   }
   `, viewable)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(f)
}

const viewable = `
query GetCustomIdFullMovie($customId: ID!) {
  viewer {
    id
    isAuthenticated
    viewableCustomId(customId: $customId) {
      ...FullViewableFields
    }
  }
}

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
`
