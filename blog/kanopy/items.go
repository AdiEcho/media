package kanopy

import (
   "net/http"
   "net/url"
)

func items() (*http.Response, error) {
   var req http.Request
   req.Header = http.Header{}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "www.kanopy.com"
   req.URL.Path = "/kapi/videos/14881163/items"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMyODIzODAzOTUzMDMxNzE5Iiwic2Vzc2lvbl9pZCI6IjE3MzI4ODU0NzUyMTAwODgxNTIiLCJjb25uZWN0aW9uX2lkIjoiMTczMjg4NTQ3NTIxMDA4ODE1MiIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzI4ODU0NzUsImV4cCI6MjA0ODI0NTQ3NSwiaXNzIjoia2FwaSJ9.rqV13JtsI1L2yyel83wA28RoOLEzu_tIkGpZ4o3yj_I"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"}
   req.Header["X-Version"] = []string{"web/prod/4.16.0/2024-11-07-14-23-23"}
   return http.DefaultClient.Do(&req)
}
