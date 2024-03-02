package spotify

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func four() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/extended-metadata/v0/extended-metadata"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body4)
   req.Header["Authorization"] = []string{"Bearer BQCkkXlvEzT-iTS4rlLwOnnAzmyxcuz7yI19Joys5qvLZxwB0XCm8bea7ikhOoioxprBD8jGa0gqnBq1wSIUXbi6Yt9iB-uZYRv5Ogwu6Ccq_59CfHlB6x8dzHeFxuvGVvQCdCQ7RMZfZ3aucXPXNNMnt_Pm8hp1dNLGeb92CKWSIf7f6UziCrBVTfJap2f0j_uHbjZamT3DKve-xhj0ViqHA30WPY6EZFhs6pzAAPmBp4hjNmheQvwMU9GWhKjvxVlJvbRV994gWlg01krDWis4CC7CsEVKOVRYBCIkg3H5vl5ymO2dNFuVvFQSCmUuWYPqx350UmulKbObUvzz"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body4 = strings.NewReader("\n\x1c\n\x02US\x12\x04free\x1a\x10\x0e\xfdN\x9d\x9c\xd8.y\x95\xd1%\xb9\xa7\x01,\xf8\x12.\n$spotify:track:1oaaSrDJimABpOdCEbw2DJ\x12\x02\b\x05\x12\x02\b\n")
