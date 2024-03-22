package peacock

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (i IdSession) Auth() (*AuthToken, error) {
   var v struct {
      Auth struct {
         AuthScheme string `json:"authScheme"`
         Proposition string `json:"proposition"`
         Provider string `json:"provider"`
         ProviderTerritory string `json:"providerTerritory"`
      } `json:"auth"`
      Device struct {
         DrmDeviceId string `json:"drmDeviceId"`
         ID string `json:"id"`
         Platform string `json:"platform"`
         Type string `json:"type"`
      } `json:"device"`
   }
   v.Auth.AuthScheme = "MESSO"
   v.Auth.Proposition = "NBCUOTT"
   v.Auth.Provider = "NBCU"
   v.Auth.ProviderTerritory = Territory
   // if empty /drm/widevine/acquirelicense will fail with
   // {
   //    "errorCode": "OVP_00306",
   //    "description": "Security failure"
   // }
   v.Device.DrmDeviceId = "UNKNOWN"
   // if incorrect /video/playouts/vod will fail with
   // {
   //    "errorCode": "OVP_00311",
   //    "description": "Unknown deviceId"
   // }
   // changing this too often will result in a four hour block
   // {
   //    "errorCode": "OVP_00014",
   //    "description": "Maximum number of streaming devices exceeded"
   // }
   v.Device.ID = "PC"
   v.Device.Platform = "ANDROIDTV"
   v.Device.Type = "TV"
   body, err := json.Marshal(v)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://ovp.peacocktv.com/auth/tokens", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(i.cookie)
   req.Header.Set("content-type", "application/vnd.tokens.v1+json")
   req.Header.Set("x-sky-signature", sign(req.Method, req.URL.Path, nil, body))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   auth := new(AuthToken)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

type IdSession struct {
   cookie *http.Cookie
}

func (i *IdSession) New(value string) {
   i.cookie = &http.Cookie{Name: "idsession", Value: value}
}

func (i IdSession) Marshal() ([]byte, error) {
   return json.MarshalIndent(i.cookie, "", " ")
}

func (i *IdSession) Unmarshal(b []byte) error {
   return json.Unmarshal(b, &i.cookie)
}
