package android

import (
   "net/http"
   "net/url"
   "os"
)

func Five() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/storage-resolve/v2/files/audio/interactive/0/f682d2a95d0e14eeef4f40b60fddde56bc6721c7"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer BQCkkXlvEzT-iTS4rlLwOnnAzmyxcuz7yI19Joys5qvLZxwB0XCm8bea7ikhOoioxprBD8jGa0gqnBq1wSIUXbi6Yt9iB-uZYRv5Ogwu6Ccq_59CfHlB6x8dzHeFxuvGVvQCdCQ7RMZfZ3aucXPXNNMnt_Pm8hp1dNLGeb92CKWSIf7f6UziCrBVTfJap2f0j_uHbjZamT3DKve-xhj0ViqHA30WPY6EZFhs6pzAAPmBp4hjNmheQvwMU9GWhKjvxVlJvbRV994gWlg01krDWis4CC7CsEVKOVRYBCIkg3H5vl5ymO2dNFuVvFQSCmUuWYPqx350UmulKbObUvzz"}
   req.URL.RawQuery = "alt=json"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
