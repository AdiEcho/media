package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

// must use IP address for correct location
func Location(content_id string) (string, error) {
   req, err := http.NewRequest("", "https://link.theplatform.com", nil)
   if err != nil {
      return "", err
   }
   req.URL.Path = func() string {
      b := []byte("/s/")
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_id...)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "assetTypes": {"DASH_CENC_HDR10|DASH_CENC_PRECON"},
      "formats": {"MPEG-DASH"},
   }.Encode()
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusFound {
      var v struct {
         Description string
      }
      json.NewDecoder(resp.Body).Decode(&v)
      return "", errors.New(v.Description)
   }
   return resp.Header.Get("location"), nil
}
