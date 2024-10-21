package max

import "net/http"

// you must
// /authentication/linkDevice/initiate
// first or this will always fail
func (b bolt_token) login() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", prd_api + "/authentication/linkDevice/login", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("cookie", "st=" + b.st)
   return http.DefaultClient.Do(req)
}
