package main

import (
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha1"
   "encoding/base64"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strings"
   "time"
)

const body = `
{"device":{"capabilities":[{"protection":"WIDEVINE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"},{"protection":"NONE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"}],"maxVideoFormat":"HD","model":"PC","hdcpEnabled":true},"client":{"thirdParties":["FREEWHEEL"]},"contentId":"GMO_00000000224510_02_HDSDR","personaParentalControlRating":"9"}
`

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play.ovp.peacocktv.com"
   req.URL.Path = "/video/playouts/vod"
   req.URL.Scheme = "https"
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
   req.Header["x-skyott-activeterritory"] = []string{"US"}
   req.Header["x-skyott-device"] = []string{"COMPUTER"}
   req.Header["x-skyott-pinoverride"] = []string{"true"}
   req.Header["x-skyott-platform"] = []string{"PC"}
   req.Header["x-skyott-proposition"] = []string{"NBCUOTT"}
   req.Header["x-skyott-provider"] = []string{"NBCU"}
   req.Header["x-skyott-territory"] = []string{"US"}
   req.Header["x-skyott-usertoken"] = []string{"13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJg311l/MwbPhAvF2BVsN57XPf+T+DHJvSb4f4vZ25jdGNdJ/fbW8YwmQInDV0Ury+V1I8/uvXLgqXQCtdQ/i23NC9RuSzTJ0LUa1Y2meoG+Vrlvy8cZSvwOxOMp6GpJB+IhZBG0iLJlYo1idT6fzD80pWPUdNM6ncp9UnlliWIh5VTXj/Fi+N6hWRgmkLshvKr0GbPVKcIY4uIV5NwslcNUAbMeI3fDaBmEfDVP7FGVM7EsayW/VbQmbu4DU5VXw5faJbINP3uDQ39LoyoH2gIcPZn7rMILVrfRgGlXabvvTDQqyTdFThChqpdVwo7rRjS0RhZGNQ3RX2CY63kKBcrJho5R/k3rj2vwIYyL++EQPHXoAnXSlUGV47JAlRq3Pi+7odT0juAtXqHuUt/Qk78RR1dehTxgzGrC5ajfl3sBgcFZD8FcZhBkFj7yvxjxaAcqA9+z5UE8ditDPSakJJxXDvVoCmH0q0yxr+DpbGWEo7JcwElv+mAoHNroezMebiQN5I/Nl3u"}
   req.Body = io.NopCloser(strings.NewReader(body))
   req.Header["X-Sky-Signature"] = []string{calculate_signature(
      req.Method, req.URL.Path, req.Header, body,
   )}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const (
   app_id = "NBCU-ANDROID-v3"
   signature_key = "JuLQgyFz9n89D9pxcN6ZWZXKWfgj2PNBUb32zybj"
   sig_version = "1.0"
)

func calculate_signature(
   method, path string, headers http.Header, payload string,
) string {
   timestamp := time.Now().Unix()
   var text_headers []string
   for key := range headers {
      if strings.HasPrefix(key, "x-skyott-") {
         text_headers = append(
            text_headers, key + ": " + headers[key][0] + "\n",
         )
      }
   }
   slices.Sort(text_headers)
   encode := strings.Join(text_headers, "")
   headers_md5 := md5.Sum([]byte(encode))
   payload_md5 := md5.Sum([]byte(payload))
   signature := func() string {
      h := hmac.New(sha1.New, []byte(signature_key))
      fmt.Fprintln(h, method)
      fmt.Fprintln(h, path)
      fmt.Fprintln(h)
      fmt.Fprintln(h, app_id)
      fmt.Fprintln(h, sig_version)
      fmt.Fprintf(h, "%x\n", headers_md5)
      fmt.Fprintln(h, timestamp)
      fmt.Fprintf(h, "%x\n", payload_md5)
      hashed := h.Sum(nil)
      return base64.StdEncoding.EncodeToString(hashed[:])
   }
   sky_ott := func() string {
      b := []byte("SkyOTT")
      b = fmt.Appendf(b, " client=%q", app_id)
      b = fmt.Appendf(b, ",signature=%q", signature())
      b = fmt.Appendf(b, `,timestamp="%v"`, timestamp)
      b = fmt.Appendf(b, ",version=%q", sig_version)
      return string(b)
   }
   return sky_ott()
}

