package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.vod-prd.s.joyn.de"
   req.URL.Path = "/v1/asset/a_p4svn4a28fq/playlist"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXlfc2lnbiI6InByb2QiLCJlbnRpdGxlbWVudF9pZCI6Ijg3MmRkMjNhLTFkNDEtNGIwZi04YzJhLTBmM2E1MDZhNTlmZiIsImNvbnRlbnRfaWQiOiJhX3A0c3ZuNGEyOGZxIiwidXNlcl9pZCI6IkpOQUEtYWUzZjgwY2YtNTc0My01MjFlLTkwOTItMDkyYTg3NzkyYzhkIiwicHJvZmlsZV9pZCI6IkpOQUEtYWUzZjgwY2YtNTc0My01MjFlLTkwOTItMDkyYTg3NzkyYzhkIiwiYW5vbnltb3VzX2lkIjoiNmJkMTZiNDgtYTQ4YS00MDBlLTljZGItMjdmMzM5MWUyZTRlIiwiY2F0YWxvZ19jb3VudHJ5IjoiREUiLCJsb2NhdGlvbl9jb3VudHJ5IjoiREUiLCJkaXN0cmlidXRpb25fdGVuYW50IjoiSk9ZTiIsImNvcHlyaWdodHMiOlsiVW5pdmVyc2FsIFN0dWRpb3MgSW5jLiBBbGwgUmlnaHRzIFJlc2VydmVkLiJdLCJqb3luX3BhY2thZ2VzIjpbIkRFX0ZSRUUiXSwiYnVzaW5lc3NfbW9kZWwiOiJBVk9EIiwicXVhbGl0eSI6IlNEIiwiYWRzX21heF9taWRyb2xsX2Jsb2NrcyI6MTAsImFkc19saW1pdF9wcmVyb2xsIjozLCJhZHNfbGltaXRfbWlkcm9sbCI6NSwiYWRzX3Rlc3QiOiIiLCJhZHNfdmFyaWFudCI6IiIsImFkc19icmVha19zcGFjaW5nIjoxMywiaWF0IjoxNzE1MTIyMTYwLCJleHAiOjE3MTUyMDg1NjB9.Y6KWmtE1Gq5BT4qiJ4W2hLyhkw6mwF8mGreLPJwBxPBV6g1LDm4Pnvn8rTtorLkBC0yZGlVtpJCpEtoXoHZthMjNQvgkCkI5JoDP2ezy-Lh5nIpXtcy9CrKJ_Y6vyvnTDSRz5PQuJpbt-CiHQP5bWxlfoUBYDvkKjMMy8okJiVHqiRoaYQ-ycBG60HhUKthaURy4EgY8v6m2QHH1ygWZJhCj7U45szwTW6Qq2YJKErKVGaecCbPrUNyJm4wC8jdMc_YMd7DhZ0KYu72tJN5FlHhbI0CnawpBfS4ivlSOfIe-yQX_yCh9ONA1-ZbU0U-AAD_h8UvBCbKLtITDmaxLwg"}
   val := make(url.Values)
   val["signature"] = []string{"3a082bf39122c422094360c39d1774897345c821"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`{"manufacturer":"unknown","platform":"browser","maxSecurityLevel":1,"streamingFormat":"dash","model":"unknown","protectionSystem":"widevine","enableDolbyAudio":false,"enableSubtitles":true,"maxResolution":1080,"variantName":"default","version":"v1"}`)
