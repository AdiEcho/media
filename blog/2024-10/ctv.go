package main

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Graphql-Client-Platform"] = []string{"entpay_web"}
   req.URL = &url.URL{}
   req.URL.Host = "www.ctv.ca"
   req.URL.Path = "/space-graphql/apq/graphql"
   value := url.Values{}
   value["operationName"] = []string{"axisContent"}
   value["extensions"] = []string{`
   {
     "persistedQuery": {
       "sha256Hash": "d6e75de9b5836cd6305c98c8d2411e336f59eb12f095a61f71d454f3fae2ecda"
     }
   }
   `}
   
   //pass
   //"id": "contentid/axis-content-2968346",
   
   value["variables"] = []string{`
   {
      "id": "contentid/axis-content-2988040",
      "subscriptions": [
         "ANIMAL_PLANET",
         "CTV",
         "CTV_COMEDY",
         "CTV_DRAMA",
         "CTV_LIFE",
         "CTV_MOVIES",
         "CTV_MTV",
         "CTV_MUCH",
         "CTV_SCIFI",
         "CTV_THROWBACK",
         "DISCOVERY",
         "DISCOVERY_SCIENCE",
         "DISCOVERY_VELOCITY",
         "E_NOW",
         "INVESTIGATION_DISCOVERY"
      ],
      "maturity": "ADULT",
      "language": "ENGLISH",
      "authenticationState": "UNAUTH",
      "playbackLanguage": "ENGLISH"
   }
   `}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   for i := 9; i >= 0; i-- {
      func() {
         resp, err := http.DefaultClient.Do(&req)
         if err != nil {
            panic(err)
         }
         defer resp.Body.Close()
         data, err := io.ReadAll(resp.Body)
         if err != nil {
            panic(err)
         }
         fmt.Println(
            i, string(data[:70]),
         )
      }()
      time.Sleep(time.Second)
   }
}
