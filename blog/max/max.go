package main

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

/*
FYI seems like you can HD with L3 on Max:
```xml
<Representation width="1280" height="692" sar="1:1" mimeType="video/mp4"
codecs="avc1.64001f" id="v7" bandwidth="5022716">
```
*/
const default_kid = "01021e5f16aa2c5ed02c550139b5ab82"

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      panic(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      panic(err)
   }
   key_id, err := hex.DecodeString(default_kid)
   if err != nil {
      panic(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, nil))
   if err != nil {
      panic(err)
   }
   key, err := module.Key(poster{}, key_id)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%x\n", key)
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

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "busy.any-any.prd.api.discomax.com"
   u.Path = "/drm-proxy/any/drm-proxy/drm/license/widevine"
   u.Scheme = "https"
   u.RawQuery = url.Values{
      "auth": {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmF0aW9uVGltZSI6IjIwMjQtMDYtMDlUMDk6Mzg6NDQuMzQ1MTE1MjQ2WiIsImVkaXRJZCI6IjE2MjNmZTRjLWVmNmUtNGRkMS1hMTBjLTRhMTgxZjVmNjU3OSIsImFwcEJ1bmRsZSI6ImJlYW0iLCJwbGF0Zm9ybSI6IndlYiIsInVzZXJJZCI6IlVTRVJJRDpib2x0OmNmMDFiNGQyLWQyMjUtNDY3OC04ZDkyLThlNDU4NTA4ZDdlOCIsInByb2ZpbGVJZCI6IlBST0ZJTEVJRDdlNjJkY2FkLTQ5NGUtNDBjMS1hZDdhLTUyZDA5YmIyNmU3MCIsImRldmljZUlkIjoiZDI0MmEzOGQtMDY5My00ZTNjLTk2MDctYWU5ZTU0OWQwMDQwIiwic3NhaSI6dHJ1ZSwic3RyZWFtVHlwZSI6InZvZCJ9.S3Fnw_qnjOH-JQvdcMoAsceHn69RKK0N8zxCNnN1yRc"},
      "keygen": {"playready"},
   }.Encode()
   return u.String(), true
}
