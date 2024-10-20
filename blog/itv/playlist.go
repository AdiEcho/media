package itv

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

//type playlist struct {
//   Playlist struct {
//      Video struct {
//         MediaFiles []struct {
//            Href string
//         }
//      }
//   }
//}

// hard geo block
func playlist() (*http.Response, error) {
   var value struct {
      Client struct {
         Id string `json:"id"`
      } `json:"client"`
      VariantAvailability struct {
         Drm         struct {
            MaxSupported string `json:"maxSupported"`
            System       string `json:"system"`
         } `json:"drm"`
         FeatureSet  []string `json:"featureset"`
         PlatformTag string   `json:"platformTag"`
      } `json:"variantAvailability"`
   }
   value.Client.Id = "browser"
   value.VariantAvailability.PlatformTag = "dotcom"
   value.VariantAvailability.Drm.MaxSupported = "L3"
   value.VariantAvailability.Drm.System = "widevine"
   // need all these to get 720:
   value.VariantAvailability.FeatureSet = []string{
      "hd",
      "mpeg-dash",
      "single-track",
      "widevine",
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "http://magni.itv.com/playlist/itvonline/ITV/10_3918_0001.001",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header["Accept"] = []string{"application/vnd.itv.vod.playlist.v4+json"}
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   return resp, nil
}
