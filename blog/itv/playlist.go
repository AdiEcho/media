package itv

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

// hard geo block
func (p *playlist) New() error {
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
   value.VariantAvailability.Drm.MaxSupported = "L3"
   value.VariantAvailability.Drm.System = "widevine"
   // need all these to get 720:
   value.VariantAvailability.FeatureSet = []string{
      "hd",
      "mpeg-dash",
      "single-track",
      "widevine",
   }
   value.VariantAvailability.PlatformTag = "dotcom"
   data, err := json.MarshalIndent(value, "", " ")
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://magni.itv.com/playlist/itvonline/ITV/10_3918_0001.001",
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   req.Header.Set("accept", "application/vnd.itv.vod.playlist.v4+json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return json.NewDecoder(resp.Body).Decode(p)
}

func (p *playlist) resolution_720() (string, bool) {
   for _, file := range p.Playlist.Video.MediaFiles {
      if file.Resolution == "720" {
         return file.Href, true
      }
   }
   return "", false
}

type playlist struct {
   Playlist struct {
      Video struct {
         MediaFiles []struct {
            Href string
            Resolution string
         }
      }
   }
}
