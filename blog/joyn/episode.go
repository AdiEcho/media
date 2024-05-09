package main

import (
   "net/http"
   "net/url"
   "os"
   "fmt"
   "io"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Joyn-Client-Version"] = []string{"5.702.6"}
   req.Header["Joyn-Country"] = []string{"DE"}
   req.Header["Joyn-Distribution-Tenant"] = []string{"JOYN"}
   req.Header["Joyn-Platform"] = []string{"web"}
   req.Header["Origin"] = []string{"https://www.joyn.de"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-site"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Header["X-Api-Key"] = []string{"4f0fd9f18abbe3cf0e87fdb556bc39c8"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.joyn.de"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   req.Method = "POST"
   body := fmt.Sprintf(`
   {
      "variables": {
         "path": "/serien/one-tree-hill/1-2-quaelende-angst"
      },
      "query": %q
   }
   `, episode_detail)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const episode_detail = `
fragment Season on Season {
  number
  id
  licenseTypes
  numberOfEpisodes
}

fragment EpisodeDetailAsset on Episode {
  __typename
  id
  title
  path
  number
  season {
      ...Season
  }
  airdate
  startsAt
  ageRating {
      minAge
      descriptorsText
  }
  licenseTypes
  description
  brands {
      path
  }
  markings
  genres {
      name
  }
  video {
      id
      duration
      type
      quality
      audioFormats {
          type
          languages {
              name
          }
      }
      videoFormats {
          type
      }
  }
  videoDescriptors {
      name
  }
  productPlacement
  heroImageDesktop: image(type: PRIMARY) {
      accentColor
      darkAccentColor: accentColor(type: DARK_VIBRANT)
      url(profile: "nextgen-web-herolandscape-1920x")
  }
  heroImageMobile: image(type: PRIMARY) {
      accentColor
      darkAccentColor: accentColor(type: DARK_VIBRANT)
      url(profile: "nextgen-webphone-primary-768x432")
  }
  series {
      id
      title
      subtype
      __typename
      path
      productionYear
      productionCompanies
      productionCountries
      copyrights
      path
      numberOfSeasons
      primaryImage: image(type: PRIMARY) {
          accentColor
          darkAccentColor: accentColor(type: DARK_VIBRANT)
          url(profile: "nextgen-webphone-primary-768x432")
      }
      heroImage: image(type: HERO_LANDSCAPE) {
          accentColor
          darkAccentColor: accentColor(type: DARK_VIBRANT)
          url(profile: "nextgen-web-herolandscape-1920x")
      }
  }
}

query EpisodeDetailPageStatic($path: String!) {
  page(path: $path) {
      __typename
      path
      ... on EpisodePage {
          episode {
              ...EpisodeDetailAsset
          }
      }
  }
}
`
