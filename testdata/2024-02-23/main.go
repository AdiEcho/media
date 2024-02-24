package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

var body = strings.NewReader(`
{"device":{"capabilities":[{"protection":"WIDEVINE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"},{"protection":"NONE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"}],"maxVideoFormat":"HD","model":"PC","hdcpEnabled":true},"client":{"thirdParties":["FREEWHEEL"]},"contentId":"GMO_00000000224510_02_HDSDR","personaParentalControlRating":"9"}
`)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"application/vnd.playvod.v1+json"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Type"] = []string{"application/vnd.playvod.v1+json"}
   req.Header["Origin"] = []string{"https://www.peacocktv.com"}
   req.Header["Referer"] = []string{"https://www.peacocktv.com/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-site"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play.ovp.peacocktv.com"
   req.URL.Path = "/video/playouts/vod"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["X-Skyott-Activeterritory"] = []string{"US"}
   req.Header["X-Skyott-Device"] = []string{"COMPUTER"}
   req.Header["X-Skyott-Pinoverride"] = []string{"true"}
   req.Header["X-Skyott-Platform"] = []string{"PC"}
   req.Header["X-Skyott-Proposition"] = []string{"NBCUOTT"}
   req.Header["X-Skyott-Provider"] = []string{"NBCU"}
   req.Header["X-Skyott-Territory"] = []string{"US"}
   req.Header["X-Skyott-Usertoken"] = []string{"13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJg311l/MwbPhAvF2BVsN57XPf+T+DHJvSb4f4vZ25jdGNdJ/fbW8YwmQInDV0Ury+V1I8/uvXLgqXQCtdQ/i23NC9RuSzTJ0LUa1Y2meoG+Vrlvy8cZSvwOxOMp6GpJB+IhZBG0iLJlYo1idT6fzD80pWPUdNM6ncp9UnlliWIh5VTXj/Fi+N6hWRgmkLshvKr0GbPVKcIY4uIV5NwslcNUAbMeI3fDaBmEfDVP7FGVM7EsayW/VbQmbu4DU5VXw5faJbINP3uDQ39LoyoH2gIcPZn7rMILVrfRgGlXabvvTDQqyTdFThChqpdVwo7rRjS0RhZGNQ3RX2CY63kKBcrJho5R/k3rj2vwIYyL++EQPHXoAnXSlUGV47JAlRq3Pi+7odT0juAtXqHuUt/Qk78RR1dehTxgzGrC5ajfl3sBgcFZD8FcZhBkFj7yvxjxaAcqA9+z5UE8ditDPSakJJxXDvVoCmH0q0yxr+DpbGWEo7JcwElv+mAoHNroezMebiQN5I/Nl3u"}
   req.Header["X-Sky-Signature"] = []string{create_signature_header()}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
