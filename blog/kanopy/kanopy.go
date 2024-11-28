package kanopy

import (
   "net/http"
   "net/url"
)

type poster struct{}

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Scheme = "https"
   u.Host = "www.kanopy.com"
   u.Path = "/kapi/licenses/widevine/1732823808506000167-0"
   return u.String(), true
}

func (poster) RequestHeader() (http.Header, error) {
   h := http.Header{
      "authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMyODIzODAzOTUzMDMxNzE5Iiwic2Vzc2lvbl9pZCI6IjE3MzI4MjM4MDM5NTMwODc0MDMiLCJjb25uZWN0aW9uX2lkIjoiMTczMjgyMzgwMzk1MzA4NzQwMyIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzI4MjM4MDMsImV4cCI6MjA0ODE4MzgwMywiaXNzIjoia2FwaSJ9.M6n5KPLzsLE1U8xuWc1tQ_gCAIUCFb4BtJXQDHd07m8"},
      "user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"},
      "x-version": {"web/prod/4.16.0/2024-11-07-14-23-23"},
   }
   return h, nil
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
