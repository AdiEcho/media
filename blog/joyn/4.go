package joyn

import (
   "encoding/json"
   "net/http"
   "bytes"
)

func (e entitlement) playlist(content_id string) (*playlist, error) {
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
      "POST", "https://api.vod-prd.s.joyn.de", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v1/asset/" + content_id + "/playlist"
   req.URL.RawQuery = "signature=" + e.signature(body)
   req.Header = http.Header{
      "authorization": {"Bearer " + e.Entitlement_Token},
      "content-type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(playlist)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (playlist) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (playlist) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type playlist struct {
   LicenseUrl string
   ManifestUrl string
}

func (p playlist) RequestUrl() (string, bool) {
   return p.LicenseUrl, true
}

func (playlist) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}
