package main

import (
   "net/http"
   "net/url"
   "os"
)

// github.com/davegonzalez/ott-boilerplate/blob/master/actions.js
var locations = []url.URL{
   { // pass
      Scheme: "https",
      Host: "api.vhx.com",
      Path: "/collections/my-dinner-with-andre/items",
      RawQuery: "site_id=59054",
   },
}

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E2NDUzNDE2ZTlmZmRjNjk4MDAxMDdhZTQ2ZWJhODg0YTU2ZDBjOGQ4NTYzMzgifQ.eyJhcHBfaWQiOjM0NDksImV4cCI6MTcxNjEwMjI5Miwibm9uY2UiOiIzMWQyNGVlN2M5NmUwYWNhIiwic2NvcGVzIjpbXSwic2Vzc2lvbl9pZCI6Iko2TFNDOWFBWUtWaGxWS2w5ZnV6T2c9PSIsInNpdGVfaWQiOjU5MDU0LCJ1c2VyX2lkIjo2NTc0NjE0Nn0.FJz_hElpDCZi2F2JFGZiiznp6KgPlUEIBRU9Rf1b4PAaep9-6e-huUl3doFaHmQD7T73SfPDXAm4OcUhmKVWJ5uLHTaPt4iMbMAATlY5v3TjhNaEs_syBWVXKejwzSbRZXBG5971YflauIfleh4fvBFNcVs6tEX1mIfMsySqpIKioIeOE5YWIYhYTrDtvr_j90VqJ2QY05O9jL8cDgOF06jP0QV78y3w1HH6E0l_UcK6Wt5Wll1LAy6Y3je5PqmYNS4tRsi2zPcujwpwjA1wgNQtQ4XfLvKoRZymZZn67CH1ZhFabPrF2hnVQpPN3-6S7vP7IARWd3CZ_NonbfPWImdzVicah79qyVMMVyaL1yPCz2IrZjv8VYKsJRIirl-xR168V7Y26uZ2xpKKHujwUgMAqj4obZsVZj17grLYBdmoq8chmPnuua1IhrFmN9AmLnPKGtL56e_uA7R4ubg61GdOwsxz6wXIAoIVnCxOVx6NmCcyTpbww3JexwhIugJP-P0wg0IseN6LZ86sFlr_er_ljrAiNfj2cslqTkPMr1b2uMa_LQZoQNttAchXIUBuaJvxNODLpeIfZ01rOy9pt-5al7VLMWjvulsSa4gUjQCbTZB2tSCupHb1768hoGDPs2Xtesck0zkY4CU4XJNfghVGaGoIXsHSQ2T0sLQzuvc"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &locations[0]
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
