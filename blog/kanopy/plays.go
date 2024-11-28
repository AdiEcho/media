package kanopy

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const x_version = "!/!/!/!"

const user_agent = "!"

func (v *video_plays) dash() (*video_manifest, bool) {
   for _, manifest := range v.Manifests {
      if manifest.ManifestType == "dash" {
         return &manifest, true
      }
   }
   return nil, false
}

type video_manifest struct {
   DrmLicenseId string
   ManifestType string
   Url string
}

type video_plays struct {
   Manifests []video_manifest
}

func (w *web_token) plays(video_id int) (*video_plays, error) {
   data, err := json.Marshal(map[string]int{
      "domainId": 2918,
      "userId": 8177465,
      "videoId": video_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/plays", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "content-type": {"application/json"},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   play := &video_plays{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}
