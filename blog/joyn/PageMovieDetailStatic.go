package main

import (
   "net/http"
   "net/url"
   "os"
   "fmt"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.joyn.de"
   req.URL.Path = "/graphql"
   val := make(url.Values)
   req.URL.Scheme = "https"
   req.Header["Joyn-Platform"] = []string{"web"}
   req.Header["X-Api-Key"] = []string{"4f0fd9f18abbe3cf0e87fdb556bc39c8"}
   val["variables"] = []string{"{\"path\":\"/filme/barry-seal-only-in-america\"}"}
   
   //val["extensions"] = []string{"{\"persistedQuery\":{\"version\":1,\"sha256Hash\":\"5cd6d962be007c782b5049ec7077dd446b334f14461423a72baf34df294d11b2\"}}"}
   
   val["extensions"] = []string{fmt.Sprintf(`{"query": %q}`, movie_detail)}
   
   // optional:
   val["operationName"] = []string{"PageMovieDetailStatic"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const movie_detail = `
fragment DetailsMovie on Movie {
  title
  ageRating {
      minAge
      descriptorsText
  }
  brands {
      path
  }
  copyrights
  description
  languages {
      code
      name
  }
  productionYear
  productionCountries
  productPlacement
  video {
      id
      duration
      audioLanguages {
          name
      }
      quality
  }
  markings
}

fragment MovieDetailAsset on Movie {
  id
  __typename
  heroImageDesktop: image(type: HERO_LANDSCAPE) {
      accentColor
      url(profile: "nextgen-web-herolandscape-1920x")
  }
  heroImageMobile: image(type: PRIMARY) {
      accentColor
      url(profile: "nextgen-webphone-primary-768x432")
  }
  primaryImage: image(type: PRIMARY) {
      id
      accentColor
      darkAccentColor: accentColor(type: DARK_VIBRANT)
      url(profile: "nextgen-web-primarycut-1920x1080")
  }
  cardImage: image(type: PRIMARY) {
      url(profile: "nextgen-web-livestill-503x283")
      urlMobile: url(profile: "nextgen-webphone-livestill-503x283")
  }
  tagline
  path
  licenseTypes
  markings
  licenseTypes
  tracking {
      agofCode
      externalAssetId
  }
  ...DetailsMovie
}

query PageMovieDetailStatic($path: String!) {
  page(path: $path) {
      ... on MoviePage {
          path
          tracking {
              pageName
              payload
          }
          id
          movie {
              ...MovieDetailAsset
          }
      }
  }
}
`
