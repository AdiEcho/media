package max

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (p *playback_request) New() {
   p.AppBundle = "beam"
   p.ApplicationSessionId = "b7804758-3377-4190-b429-ea7dee273880"
   p.ConsumptionType = "streaming"
   p.DeviceInfo.Player.MediaEngine.Version = "2.15.1"
   p.DeviceInfo.Player.MediaEngine.Version = "GLUON_BROWSER"
   p.DeviceInfo.Player.PlayerView.Height = 864
   p.DeviceInfo.Player.PlayerView.Width = 1536
   p.DeviceInfo.Player.Sdk.Name = "Beam Player Desktop"
   p.DeviceInfo.Player.Sdk.Version = "4.1.0"
   p.EditId = "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579"
   p.FirstPlay = true
   p.PlaybackSessionId = "6ccb51f6-5cd6-4ce0-9ade-7f8a47e66474"
}

func (r login_response) playback(p playback_request) (*http.Response, error) {
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
   
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi1iODRmOTIxMy0yMzA0LTQ4MGEtOGEzMy1lOTViMTNhOGFiZjgiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6Y2YwMWI0ZDItZDIyNS00Njc4LThkOTItOGU0NTg1MDhkN2U4IiwiaWF0IjoxNzE3OTcwODYxLCJleHAiOjIwMzMzMzA4NjEsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6ZmFsc2UsImRldmljZUlkIjoiMDQxMzdhYTItMWUxZS02ZjUyLTdhMDgtMTIyNDljODY0NjkwIn0.adU124rWw6-B55slVSnAn7gyd6wJA8sdWv-c2ayXkdrlGmXSRIosAnxf582ABO2ZCmguG0Lbm2S2ZlKMuRSwdT-QXfG8-EFW4LAaawiMc3xKuRn-uUmMCAhaewg_4TauEFdpAPDXAFOdO_wItNt7MoN1nQaW8C1Sa7jJzDpQhCDqv8DfEeZYfx_jQopnVyw6vUmz_W4m52wJAlmh_kW2fCJuUahywMKRHSBOBriBm1LL51gIOIcFxfLM3G6f-yb_ar9xqFqSIyaSpQ5Wj1t2T3IQPNh8gjrPw0O6A2zGo91S4reQHDB18kc-IzCFfGA2scGbLWZE7rcUpeZjK6eRiQ",
   }
   return http.DefaultClient.Do(&req)
}

type playback_request struct {
   AppBundle            string `json:"appBundle"`
   ApplicationSessionId string `json:"applicationSessionId"`
   Capabilities         struct {
      Manifests struct {
         Formats struct {
            Dash struct {
            } `json:"dash"`
         } `json:"formats"`
      } `json:"manifests"`
   } `json:"capabilities"`
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo      struct {
      Player struct {
         MediaEngine struct {
            Name    string `json:"name"`
            Version string `json:"version"`
         } `json:"mediaEngine"`
         PlayerView struct {
            Height int `json:"height"`
            Width  int `json:"width"`
         } `json:"playerView"`
         Sdk struct {
            Name    string `json:"name"`
            Version string `json:"version"`
         } `json:"sdk"`
      } `json:"player"`
   } `json:"deviceInfo"`
   EditId            string `json:"editId"`
   FirstPlay         bool   `json:"firstPlay"`
   Gdpr              bool   `json:"gdpr"`
   PlaybackSessionId string `json:"playbackSessionId"`
   UserPreferences   struct{} `json:"userPreferences"`
}
