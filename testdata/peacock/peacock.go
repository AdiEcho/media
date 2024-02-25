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

const (
   app_id = "NBCU-ANDROID-v3"
   signature_key = "JuLQgyFz9n89D9pxcN6ZWZXKWfgj2PNBUb32zybj"
   sig_version = "1.0"
)

func (a auth_tokens) video(content_id string) (*video_playouts, error) {
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
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://play.ovp.peacocktv.com/video/playouts/vod",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("x-skyott-usertoken", a.UserToken)
   // `application/json` fails
   req.Header.Set("content-type", "application/vnd.playvod.v1+json")
   req.Header.Set(
      "x-sky-signature", sign(req.Method, req.URL.Path, req.Header, body),
   )
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   video := new(video_playouts)
   if err := json.NewDecoder(res.Body).Decode(video); err != nil {
      return nil, err
   }
   return video, nil
}

type video_playouts struct {
   Asset struct {
      Endpoints []struct {
         CDN string
         URL string
      }
   }
   Protection struct {
      LicenceAcquisitionUrl string // wikipedia.org/wiki/License
   }
}

func (v video_playouts) RequestUrl() (string, bool) {
   return v.Protection.LicenceAcquisitionUrl, true
}

func (video_playouts) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (video_playouts) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
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

// how the fuck do we sign the body?
func (video_playouts) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set(
      "x-sky-signature",
      sign("POST", "/drm/widevine/acquirelicense", nil, nil),
   )
   return h, nil
}
