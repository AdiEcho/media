package max

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (st st_cookie) playback(p playback_request) (*http.Response, error) {
   body, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.max.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/any/playback/v1/playbackInfo"
   req.AddCookie(st.Cookie)
   req.Header.Set("content-type", "application/json")
   return http.DefaultClient.Do(req)
}

func (p *playback_request) New() {
   p.ConsumptionType = "streaming"
   p.EditId = "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579"
}

type playback_request struct {
   AppBundle string `json:"appBundle"` // required
   ApplicationSessionId string `json:"applicationSessionId"` // required
   Capabilities struct {
      Manifests struct {
         Formats struct {
            Dash struct{} `json:"dash"` // required
         } `json:"formats"` // required
      } `json:"manifests"` // required
   } `json:"capabilities"` // required
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo struct {
      Player struct {
         MediaEngine struct {
            Name string `json:"name"` // required
            Version string `json:"version"` // required
         } `json:"mediaEngine"` // required
         PlayerView struct {
            Height int `json:"height"` // required
            Width int `json:"width"` // required
         } `json:"playerView"` // required
         Sdk struct {
            Name string `json:"name"` // required
            Version string `json:"version"` // required
         } `json:"sdk"` // required
      } `json:"player"` // required
   } `json:"deviceInfo"` // required
   EditId string `json:"editId"`
   FirstPlay bool `json:"firstPlay"` // required
   Gdpr bool `json:"gdpr"` // required
   PlaybackSessionId string `json:"playbackSessionId"` // required
   UserPreferences struct{} `json:"userPreferences"` // required
}
