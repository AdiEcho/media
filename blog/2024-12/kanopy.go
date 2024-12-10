package main

import (
   "net/http"
   "net/url"
   "os"
)

/*
show
kanopy.com/en/product/14881161

season
kanopy.com/en/product/14881163

episode
kanopy.com/en/product/14881167

"Wildfire", 1, 2, "The Rescue"

"Wildfire", 14881167, "The Rescue"
*/
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMzNTg5NjM1Mjc0MDI3MTUxIiwic2Vzc2lvbl9pZCI6IjE3MzM1ODk2MzUyNzQwNDU5ODMiLCJjb25uZWN0aW9uX2lkIjoiMTczMzU4OTYzNTI3NDA0NTk4MyIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzM1ODk2MzUsImV4cCI6MjA0ODk0OTYzNSwiaXNzIjoia2FwaSJ9.sTWKCUoZ2APHMtD2zXQ0nb5pt-dTZu6YyXfbidMniY4"}
   req.Header["User-Agent"] = []string{"!"}
   req.Header["X-Version"] = []string{"!/!/!/!"}
   req.URL = &url.URL{}
   req.URL.Host = "www.kanopy.com"
   
   // show
   // req.URL.Path = "/kapi/videos/14881161"
   
   // title
   req.URL.Path = "/kapi/videos/14881167"
   
   req.URL.Scheme = "https"
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
