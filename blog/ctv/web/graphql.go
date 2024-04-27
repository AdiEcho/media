package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   // you need this for the first request, then can omit
   req.Header["Graphql-Client-Platform"] = []string{"entpay_web"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.ctv.ca"
   req.URL.Path = "/space-graphql/apq/graphql"
   val := make(url.Values)
   val["operationName"] = []string{"resolvePath"}
   
   val["variables"] = []string{`
   {
     "page": 0,
     "path": "/movies/the-girl-with-the-dragon-tattoo-2011/the-girl-with-the-dragon-tattoo-2011",
     "subscriptions": [
       "CTV",
       "CTV_DRAMA",
       "CTV_COMEDY",
       "CTV_LIFE",
       "CTV_SCIFI",
       "CTV_THROWBACK",
       "CTV_MOVIES",
       "CTV_MTV",
       "CTV_MUCH",
       "DISCOVERY",
       "DISCOVERY_SCIENCE",
       "DISCOVERY_VELOCITY",
       "INVESTIGATION_DISCOVERY",
       "ANIMAL_PLANET",
       "E_NOW"
     ],
     "maturity": "ADULT",
     "language": "ENGLISH",
     "authenticationState": "UNAUTH",
     "playbackLanguage": "ENGLISH"
   }
   `}
   val["extensions"] = []string{`
   {
      "persistedQuery": {
         "version": 1,
         "sha256Hash": "26d314b59ba2708d261067964353f9a92f1c2689f50d1254fa4d03ddb9b9092a"
      }
   }
   `}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
