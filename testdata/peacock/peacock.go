package peacock

import (
   "bytes"
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha1"
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "time"
)

func (v video_playouts) sign(method, path string, body []byte) string {
   timestamp := time.Now().Unix()
   encode := func() []byte {
      b := []byte("x-skyott-usertoken: ")
      b = append(b, v.user_token()...)
      return append(b, '\n')
   }
   headers_md5 := md5.Sum(encode())
   payload_md5 := md5.Sum(body)
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
      // must be quoted
      b = fmt.Appendf(b, " client=%q", app_id)
      // must be quoted
      b = fmt.Appendf(b, ",signature=%q", signature())
      // must be quoted
      b = fmt.Appendf(b, `,timestamp="%v"`, timestamp)
      // must be quoted
      b = fmt.Appendf(b, ",version=%q", sig_version)
      return string(b)
   }
   return sky_ott()
}

type video_playouts struct {
   Protection struct {
      LicenceToken string // wikipedia.org/wiki/License
   }
}

const (
   app_id = "NBCU-ANDROID-v3"
   signature_key = "JuLQgyFz9n89D9pxcN6ZWZXKWfgj2PNBUb32zybj"
   sig_version = "1.0"
)

func (v *video_playouts) New(content_id string) error {
   body, err := func() ([]byte, error) {
      type capability struct {
         Acodec string `json:"acodec"`
         Container string `json:"container"`
         Protection string `json:"protection"`
         Transport string `json:"transport"`
         Vcodec string `json:"vcodec"`
      }
      var s struct {
         ContentId string `json:"contentId"`
         Device struct {
            Capabilities []capability `json:"capabilities"`
         } `json:"device"`
      }
      s.ContentId = content_id
      s.Device.Capabilities = []capability{
         {
            Acodec: "AAC",
            Container: "ISOBMFF",
            Protection: "WIDEVINE",
            Transport: "DASH",
            Vcodec: "H264",
         },
      }
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://play.ovp.peacocktv.com/video/playouts/vod",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   // `application/json` fails
   req.Header.Set("content-type", "application/vnd.playvod.v1+json")
   req.Header.Set("x-skyott-usertoken", v.user_token())
   req.Header.Set("x-sky-signature", v.sign(req.Method, req.URL.Path, body))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(res.Body).Decode(v)
}

func (video_playouts) user_token() string {
   return "13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJg311l/MwbPhAvF2BVsN57XPf+T+DHJvSb4f4vZ25jdGNdJ/fbW8YwmQInDV0Ury+V1I8/uvXLgqXQCtdQ/i23NC9RuSzTJ0LUa1Y2meoG+Vrlvy8cZSvwOxOMp6GpJB+IhZBG0iLJlYo1idT6fzD80pWPUdNM6ncp9UnlliWIh5VTXj/Fi+N6hWRgmkLshvKr0GbPVKcIY4uIV5NwslcNUAbMeI3fDaBmEfDVP7FGVM7EsayW/VbQmbu4DU5VXw5faJbINP3uDQ39LoyoH2gIcPZn7rMILVrfRgGlXabvvTDQqyTdFThChqpdVwo7rRjS0RhZGNQ3RX2CY63kKBcrJho5R/k3rj2vwIYyL++EQPHXoAnXSlUGV47JAlRq3Pi+7odT0juAtXqHuUt/Qk78RR1dehTxgzGrC5ajfl3sBgcFZD8FcZhBkFj7yvxjxaAcqA9+z5UE8ditDPSakJJxXDvVoCmH0q0yxr+DpbGWEo7JcwElv+mAoHNroezMebiQN5I/Nl3u"
}
