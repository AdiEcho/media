package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.Header["Magine-Play-Devicemodel"] = []string{"firefox 111.0 / windows 10"}
   req.Header["Magine-Play-Deviceplatform"] = []string{"firefox"}
   req.Header["Magine-Play-Devicetype"] = []string{"web"}
   req.Header["Magine-Play-Drm"] = []string{"widevine"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL.Host = "client-api.magine.com"
   req.URL.Scheme = "https"
   req.Header["Magine-Play-Deviceid"] = []string{"!"}
   req.Header["Magine-Play-Protocol"] = []string{"dashs"}
   req.Header["x-forwarded-for"] = []string{"78.64.0.0"}
   req.URL.Path = "/api/playback/v1/preflight/asset/8149455c-cb3d-4b15-85a8-b95e3d1570b5"
   req.Header["Magine-Accesstoken"] = []string{"22cc71a2-8b77-4819-95b0-8c90f4cf5663"}
   req.Header["Magine-Play-Entitlementid"] = []string{"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJhbGxvd2VkQ291bnRyaWVzIjpbIklFIiwiRlIiLCJOTCIsIlNJIiwiUk8iLCJCRyIsIkNZIiwiSVMiLCJERSIsIkZJIiwiTFYiLCJQTCIsIlBUIiwiTFUiLCJIUiIsIkVTIiwiQVQiLCJNVCIsIlNFIiwiR1IiLCJJVCIsIkhVIiwiRUUiLCJESyIsIkxJIiwiTFQiLCJTSyIsIkNaIiwiQkUiLCJOTyJdLCJtYXJrZXRJZHMiOlsiU0UiXSwiZXhwIjoxNzE0OTA1NTMwLCJhZHMiOmZhbHNlLCJpYXQiOjE3MTQ4NjIzMzAsInN1YiI6IjE1NktFRkpETEM3SE1JS0FVTlNRQkc4TFNVU1IiLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaXNzIjoiZHJha2VuZmlsbSIsIm9mZmxpbmVFeHBpcmF0aW9uIjoxNzE1MzEyNTAyLCJ1c2VyVGFncyI6W10sIm9mZmVySWQiOiI3OTNGTzg4NFFZVUJXRVM0TTEyRk1ESThKT0ZSIiwiYXNzZXRJZCI6IjgxNDk0NTVjLWNiM2QtNGIxNS04NWE4LWI5NWUzZDE1NzBiNSJ9.tNkc_ZsE2j1cZwNLg0wqpucU3HJcmvwWd0sqVlqijH5bhVEVON9Xmf5mBlb_UQLLhUW_3mMCTvpOfyn38yZclQ"}
   //////////////////////////////////////////////////////////////////////////////
   req.Header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODYyMzI3LCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.sYyyMBA7gf0q7A9na8E-vkgJntedFYn2pk_LX2WYBgdQgLgNs7xrtUgR2ZoZlMhgN6D5rQj2U6WDzvDUHZCqEQ"}
   //////////////////////////////////////////////////////////////////////////////
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
