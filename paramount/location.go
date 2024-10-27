package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

// hard geo block
// 
// AU:
// formats=MPEG-DASH&assetTypes=DASH_CENC
// FR:
// formats=MPEG-DASH&assetTypes=DASH_CENC_PRECON
// US:
// formats=MPEG-DASH&assetTypes=DASH_CENC
func Location(content_id, asset_type string) (string, error) {
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
      "assetTypes": {asset_type},
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
