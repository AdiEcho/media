package peacock

import (
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha1"
   "encoding/base64"
   "fmt"
   "net/http"
   "slices"
   "strings"
   "time"
)

const body = `
{"device":{"capabilities":[{"protection":"WIDEVINE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"},{"protection":"NONE","container":"ISOBMFF","transport":"DASH","acodec":"AAC","vcodec":"H264"}],"maxVideoFormat":"HD","model":"PC","hdcpEnabled":true},"client":{"thirdParties":["FREEWHEEL"]},"contentId":"GMO_00000000224510_02_HDSDR","personaParentalControlRating":"9"}
`

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

