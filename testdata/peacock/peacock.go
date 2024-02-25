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
   "slices"
   "strings"
   "time"
)

func user_token() string {
   return "38-uyyWV0cL0it8NIN0igaMbD045ycSbGKUl9kuw5KyZuQIPgBSpm32E5bzXTFmysyIsgKawlGqKccNLHhor1ru5gEbU6rL/FfTSTTnNW70XO7TyHVAwKC0AhQEh0R4TSryjRAZjX06UMMJltKccv3pjTlsW52Y5Wo3QMjURNPZ2JA5yvjWhl2e5E+ZNpFGnhUj2YhRgj+Pjrb9b5hJhev+iSxTaDXcROivICKPdqxAlMxP43POUPhuxmzeNJ8ZtyWoGpFHQjjetpLPTKubJ6eyLA4V4/MGzNvt4gA4B7BbbWBw05tHSFno3Sgaxp9jSZv9XiNoWkhHHV25iVZyU1fcAfCkGzDctFQFFYL3D0MI9NWaeqKBD2a9cngxdm5Sd/eEm/0mJeeQ7Hllhhg0WELn/S7Q+uiv7/mCPBtW+tI6nlsvDSZyaouylM/8cxXQnZ1z5WP4gpprWAdMcQBFwzMVLYv9TvN09uZNWxVKXIox7nA0sATaN82nUbuq0P3hrWMtaEabOrYQqEHk0yN1ThPStJ4G50aRVGVFbbyUiAAgRijWPySbeuPQJ9612BWgvEtd"
}

func sign(method, path string, head http.Header, body []byte) string {
   timestamp := time.Now().Unix()
   text_headers := func() string {
      var s []string
      for k := range head {
         k = strings.ToLower(k)
         if strings.HasPrefix(k, "x-skyott-") {
            s = append(s, k + ": " + head.Get(k) + "\n")
         }
      }
      slices.Sort(s)
      return strings.Join(s, "")
   }()
   headers_md5 := md5.Sum([]byte(text_headers))
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
   req.Header.Set("x-skyott-usertoken", user_token())
   req.Header.Set(
      "x-sky-signature", sign(req.Method, req.URL.Path, req.Header, body),
   )
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
