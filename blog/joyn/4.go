package joyn

import (
   "encoding/json"
   "net/http"
   "bytes"
)

func (e entitlement) playlist(m movie_detail) (*http.Response, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Manufacturer     string `json:"manufacturer"`
         MaxResolution    int    `json:"maxResolution"`
         Model            string `json:"model"`
         Platform         string `json:"platform"`
         ProtectionSystem string `json:"protectionSystem"`
         StreamingFormat  string `json:"streamingFormat"`
      }
      s.Manufacturer = "unknown"
      s.MaxResolution = 1080
      s.Model = "unknown"
      s.Platform = "browser"
      s.ProtectionSystem = "widevine"
      s.StreamingFormat = "dash"
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.vod-prd.s.joyn.de/v1/asset/a_p4svn4a28fq/playlist",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + e.Entitlement_Token},
      "content-type": {"application/json"},
   }
   req.URL.RawQuery = "signature=" + e.signature(body)
   return http.DefaultClient.Do(req)
}
