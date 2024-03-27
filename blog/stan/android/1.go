package stan

import (
   "encoding/json"
   "io"
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
   b.WriteString(a.v.Code)
   return b.String()
}

type activation_code struct {
   data []byte
   v struct {
      Code string
      URL string
   }
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
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *activation_code) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
