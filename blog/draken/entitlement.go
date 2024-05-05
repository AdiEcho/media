package draken

import (
   "encoding/json"
   "errors"
   "net/http"
)

type entitlement struct {
   Token string
}

func (a auth_login) entitlement(f *full_movie) (*entitlement, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/entitlement/v2/asset/" + f.ID
   req.Header = http.Header{
      "authorization": {"Bearer " + a.v.Token},
      "magine-accesstoken": {magine_accesstoken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   title := new(entitlement)
   err = json.NewDecoder(res.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}
