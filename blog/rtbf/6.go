package rtbf

import (
   "net/http"
   "strings"
)

func (g gigya_login) six() (*http.Response, error) {
   req, err := http.NewRequest("", "https://exposure.api.redbee.live", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement")
      b.WriteString("/3201987_6BA97Bb/play")
      return b.String()
   }()
   req.Header = http.Header{
      "x-forwarded-for": {"91.90.123.17"},
      "authorization": {"Bearer " + g.SessionToken},
   }
   return http.DefaultClient.Do(req)
}
