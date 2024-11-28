package kanopy

import (
   "net/http"
   "net/url"
)

func (web_token) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (web_token) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (web_token) RequestUrl() (string, bool) {
   var u url.URL
   u.Scheme = "https"
   u.Host = "www.kanopy.com"
   u.Path = "/kapi/licenses/widevine/1732823808506000167-0"
   return u.String(), true
}

func (w *web_token) RequestHeader() (http.Header, error) {
   h := http.Header{
      //"authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMyODIzODAzOTUzMDMxNzE5Iiwic2Vzc2lvbl9pZCI6IjE3MzI4MjM4MDM5NTMwODc0MDMiLCJjb25uZWN0aW9uX2lkIjoiMTczMjgyMzgwMzk1MzA4NzQwMyIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzI4MjM4MDMsImV4cCI6MjA0ODE4MzgwMywiaXNzIjoia2FwaSJ9.M6n5KPLzsLE1U8xuWc1tQ_gCAIUCFb4BtJXQDHd07m8"},
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {"!"},
      "x-version": {"!/!/!/!"},
   }
   return h, nil
}
