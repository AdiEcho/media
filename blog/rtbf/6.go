package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

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
