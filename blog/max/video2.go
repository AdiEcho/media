package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "default.any-amer.prd.api.max.com"
   req.URL.Path = "/cms/routes/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["include"] = []string{"default"}
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi03ZTgyMDViMS0wOGZmLTQyYzEtYjZmMC0yNTdjMmVmZWRkMTMiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6Y2YwMWI0ZDItZDIyNS00Njc4LThkOTItOGU0NTg1MDhkN2U4IiwiaWF0IjoxNzE4MzE4MDIyLCJleHAiOjIwMzM2NzgwMjIsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6ZmFsc2UsImRldmljZUlkIjoiMzQ1ZGNjMjYtN2U5Yy00Y2JlLTk1ODEtYTk1ZmQ5NDI3YTdiIn0.GY7OoIUcSKKtN-dnCqAQkHZY7rZU55t5wQQRCx6BROmD-DXCmbPVsUrSHkcU4s0k3vXyTqEjGrI8cCCEtWe60-QFD_Fh7FFykkPb81jqBaEYNLEvBqwb5Un6nsXCcwhcF9bp4sj-uh8EcI9M12SpORH-nQwhOOUi340VjonT5aJ0bgNsdWv56v37U83h0u2BEoqShG2gDIAZuR1Aaqd9XHD5HeFTDVbYPga9MRkBf8hcU5lIFSrSR66D4f__A-0WXAlcm1EnWaCXAaI1JQxa4mrhh-Ulp_0nWQQ_x2StGR4FsA34tVPlGSROsSwSQTPkXkMuwMKvAI5NQBkeZzj_bA",
   }
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
