package max

import "net/http"

type default_routes struct {
   Included []struct {
      Attributes struct {
         Type string
      }
   }
}

func (d default_token) route_android(path string) (*http.Response, error) {
   req, err := http.NewRequest(
      "", "https://default.any-amer.prd.api.discomax.com/cms/routes"+path, nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "include=default"
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   return http.DefaultClient.Do(req)
}
