package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (p Playback) RequestUrl() (string, error) {
   return p.DRM.Widevine.LicenseServer, nil
}

func (c CrossSite) Playback(id string) (*Playback, error) {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "mediaFormat": "mpeg-dash",
         "providerId": "rokuavod",
         "rokuId": id,
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://therokuchannel.roku.com/api/v3/playback",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   // we could use Request.AddCookie, but we would need to call it after this,
   // otherwise it would be clobbered
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
      "Cookie": {c.csrf().Raw},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (Playback) RequestHeader() http.Header {
   return nil
}

func (Playback) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type CrossSite struct {
   cookies []*http.Cookie
   token string
}

func (c CrossSite) csrf() *http.Cookie {
   for _, cookie := range c.cookies {
      if cookie.Name == "_csrf" {
         return cookie
      }
   }
   return nil
}

type MediaVideo struct {
   DrmAuthentication *struct{}
   URL string
   VideoType string
}
