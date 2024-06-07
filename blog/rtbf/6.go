package rtbf

import (
   "encoding/json"
   "net/http"
   "strings"
)

type entitlement struct {
   AssetId string
   PlayToken string
}

func (g gigya_login) entitlement() (*entitlement, error) {
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   title := new(entitlement)
   err = json.NewDecoder(res.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}
