package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"}
   req.Header["X-Requested-With"] = []string{"com.roku.web.trc"}
   req.Header["X-Roku-Code-Version"] = []string{"2"}
   req.Header["X-Roku-Content-Token"] = []string{"ZizZafDEidXT//m06oXf0V7G0vSINNHxlL/mCTNv2TVL89l4++Su+yNDIGA1bD5CWEmcxaYfcm736PCwFKPsriM0vuRi6gW8+HXVnz847RVoGju10f599ErmxKFtBF9u8tA6KiGbKVE5RymAmP1cW4ae9fgsQU9JqB3QtHPwdqSZgPNM3Ue9+nqkJZt+EV2MGKDy2kHi17ggFDaylNTcsDRmauBFrGdP8RhPaW9uL3FWPGBZ5P2wIBGmNSWsKiToXMes1D/OoDlK76G9SegdrGScdR+7QzZBCNK2I4D/y2ijLX2CIwFoNoVmCrU8PUg3mt81GXFxLbOpszAHOYPHFCtlKikq1bQLjY/0yKQ6N8PgEi63ZXd3a5ekjTVFq+l1E4DBmyBUEx6E21jOS11Ahg=="}
   req.Header["X-Roku-Reserved-Channel-Store-Code"] = []string{"us"}
   req.Header["X-Roku-Reserved-Culture-Code"] = []string{"en-US"}
   req.Header["X-Roku-Reserved-Experiment-Configs"] = []string{"e30="}
   req.Header["X-Roku-Reserved-Experiment-State"] = []string{"W10="}
   req.Header["X-Roku-Reserved-Session-Id"] = []string{"f77813e1-a689-41e9-b058-097e8520f4d2"}
   req.Header["X-Roku-Reserved-Time-Zone-Offset"] = []string{"+00:00"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "googletv.web.roku.com"
   req.URL.Path = "/api/v1/account/activation/l8pkog"
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
