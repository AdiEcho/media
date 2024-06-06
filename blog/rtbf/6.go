package rtbf

import (
   "net/http"
   "net/url"
   "os"
)

func six() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "exposure.api.redbee.live"
   req.URL.Scheme = "https"
   req.Header.Set("x-forwarded-for", "91.90.123.17")
   req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/entitlement/3201987_6BA97Bb/play"
   req.Header["Authorization"] = []string{"Bearer ses_c25be097-b9e2-4bdc-b12f-e70a3556b910p|acc_8982699050ed41f98fed0877bda6f616_6BA97Bb|usr_8982699050ed41f98fed0877bda6f616_6BA97Bb|null|1717634395676|2017634395638|false|7f5cdd55-1cfe-4841-9e8e-ecd8b823cfad|WEB||RTBFAuvio||2iyx75pBySG7rz7wJh6HW4NrObpoUN+enZLRTJzkzJw="}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
