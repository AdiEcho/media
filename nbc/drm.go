package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

type DrmProxy struct {
   Hash string
   Time string
}

const drm_proxy_secret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"

func (d *DrmProxy) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://drmproxy.digitalsvc.apps.nbcuni.com",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/drm-proxy/license/widevine"
   req.URL.RawQuery = url.Values{
      "device": {"web"},
      "hash": {d.Hash},
      "time": {d.Time},
   }.Encode()
   req.Header.Set("content-type", "application/octet-stream")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (d *DrmProxy) New() {
   d.Time = fmt.Sprint(time.Now().UnixMilli())
   d.Hash = func() string {
      hash := hmac.New(sha256.New, []byte(drm_proxy_secret))
      hash.Write([]byte(d.Time))
      hash.Write([]byte("widevine"))
      return fmt.Sprintf("%x", hash.Sum(nil))
   }()
}

