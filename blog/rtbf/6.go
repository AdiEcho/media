package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (entitlement) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set("content-type", "application/x-protobuf")
   return h, nil
}

func (entitlement) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (entitlement) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (g gigya_login) entitlement(embed embed_media) (*entitlement, error) {
   req, err := http.NewRequest("", "https://exposure.api.redbee.live", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement/")
      b.WriteString(embed.Data.AssetId)
      b.WriteString("/play")
      return b.String()
   }()
   req.Header = http.Header{
      "x-forwarded-for": {"91.90.123.17"},
      "authorization": {"Bearer " + g.SessionToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   title := new(entitlement)
   err = json.NewDecoder(res.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

type entitlement struct {
   AssetId string
   PlayToken string
}

func (e entitlement) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "rbm-rtbf.live.ott.irdeto.com"
   u.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
   u.Scheme = "https"
   u.RawQuery = url.Values{
      "contentId": {e.AssetId},
      "ls_session": {e.PlayToken},
   }.Encode()
   return u.String(), true
}
