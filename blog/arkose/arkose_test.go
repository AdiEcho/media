package arkose

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "testing"
)

func TestBda(t *testing.T) {
   bda, err := get_bda()
   if err != nil {
      t.Fatal(err)
   }
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "wbd-api.arkoselabs.com"
   req.URL.Path = "/fc/gt2/public_key/B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   data := url.Values{
      "public_key":[]string{"B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"},
      "bda":[]string{bda},
   }.Encode()
   req.Body = io.NopCloser(strings.NewReader(data))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
