package mubi

import "net/http"

func films(path string) (*http.Response, error) {
   req, err := http.NewRequest("GET", "https://api.mubi.com/v3" + path, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   return http.DefaultClient.Do(req)
}
