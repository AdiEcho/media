package max

import (
   "net/http"
   "net/url"
)

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "busy.any-any.prd.api.discomax.com"
   u.Path = "/drm-proxy/any/drm-proxy/drm/license/widevine"
   u.Scheme = "https"
   u.RawQuery = url.Values{
      "keygen": {"playready"},
      "auth": {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmF0aW9uVGltZSI6IjIwMjQtMDYtMDlUMDk6Mzg6NDQuMzQ1MTE1MjQ2WiIsImVkaXRJZCI6IjE2MjNmZTRjLWVmNmUtNGRkMS1hMTBjLTRhMTgxZjVmNjU3OSIsImFwcEJ1bmRsZSI6ImJlYW0iLCJwbGF0Zm9ybSI6IndlYiIsInVzZXJJZCI6IlVTRVJJRDpib2x0OmNmMDFiNGQyLWQyMjUtNDY3OC04ZDkyLThlNDU4NTA4ZDdlOCIsInByb2ZpbGVJZCI6IlBST0ZJTEVJRDdlNjJkY2FkLTQ5NGUtNDBjMS1hZDdhLTUyZDA5YmIyNmU3MCIsImRldmljZUlkIjoiZDI0MmEzOGQtMDY5My00ZTNjLTk2MDctYWU5ZTU0OWQwMDQwIiwic3NhaSI6dHJ1ZSwic3RyZWFtVHlwZSI6InZvZCJ9.S3Fnw_qnjOH-JQvdcMoAsceHn69RKK0N8zxCNnN1yRc"},
   }.Encode()
   return u.String(), true
}

type poster struct{}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}
