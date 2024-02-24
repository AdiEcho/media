package peacock

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "testing"
)

func TestPeacock(t *testing.T) {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play.ovp.peacocktv.com"
   req.URL.Path = "/video/playouts/vod"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/vnd.playvod.v1+json"}
   req.Header["x-skyott-usertoken"] = []string{"13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJg311l/MwbPhAvF2BVsN57XPf+T+DHJvSb4f4vZ25jdGNdJ/fbW8YwmQInDV0Ury+V1I8/uvXLgqXQCtdQ/i23NC9RuSzTJ0LUa1Y2meoG+Vrlvy8cZSvwOxOMp6GpJB+IhZBG0iLJlYo1idT6fzD80pWPUdNM6ncp9UnlliWIh5VTXj/Fi+N6hWRgmkLshvKr0GbPVKcIY4uIV5NwslcNUAbMeI3fDaBmEfDVP7FGVM7EsayW/VbQmbu4DU5VXw5faJbINP3uDQ39LoyoH2gIcPZn7rMILVrfRgGlXabvvTDQqyTdFThChqpdVwo7rRjS0RhZGNQ3RX2CY63kKBcrJho5R/k3rj2vwIYyL++EQPHXoAnXSlUGV47JAlRq3Pi+7odT0juAtXqHuUt/Qk78RR1dehTxgzGrC5ajfl3sBgcFZD8FcZhBkFj7yvxjxaAcqA9+z5UE8ditDPSakJJxXDvVoCmH0q0yxr+DpbGWEo7JcwElv+mAoHNroezMebiQN5I/Nl3u"}
   req.Body = io.NopCloser(strings.NewReader(body))
   req.Header["X-Sky-Signature"] = []string{calculate_signature(
      req.Method, req.URL.Path, req.Header, body,
   )}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
