package member

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type authenticate struct {
   Data struct {
      UserAuthenticate struct {
         Access_Token string
      }
   }
}

func (a *authenticate) New(email, password string) error {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.audienceplayer.com"
   req.URL.Path = "/graphql/2/user"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   body := fmt.Sprintf(`
   {
      "variables": {
         "email": %q,
         "password": %q
      },
      "query": %q
   }
   `, email, password, user_authenticate)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

const user_authenticate = `
mutation UserAuthenticate($email: String, $password: String) {
  UserAuthenticate(email: $email, password: $password) {
    access_token
  }
}
`
