package stan

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func (a activation_code) String() string {
   var b strings.Builder
   b.WriteString("Stan.\n")
   b.WriteString("Log in with code\n")
   b.WriteString("1. Visit stan.com.au/activate\n")
   b.WriteString("2. Enter the code:\n")
   b.WriteString(a.Code)
   return b.String()
}

type activation_code struct {
   Code string
   Token string
   URL string
}

func (a *activation_code) New() error {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/activation-codes/", url.Values{
         "generate": {"true"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
