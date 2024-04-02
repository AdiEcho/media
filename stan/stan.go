package stan

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

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

func (a activation_code) String() string {
   var b strings.Builder
   b.WriteString("Stan.\n")
   b.WriteString("Log in with code\n")
   b.WriteString("1. Visit stan.com.au/activate\n")
   b.WriteString("2. Enter the code:\n")
   b.WriteString(a.v.Code)
   return b.String()
}

func (a activation_code) token() (*web_token, error) {
   res, err := http.Get(a.v.URL)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var web web_token
   web.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &web, nil
}

func (a *activation_code) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}

type app_session struct {
   JwToken string
}

type web_token struct {
   data []byte
   v struct {
      JwToken string
      ProfileId string
   }
}

func (w web_token) session() (*app_session, error) {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/sessions/mobile/app", url.Values{
         "jwToken": {w.v.JwToken},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   session := new(app_session)
   if err := json.NewDecoder(res.Body).Decode(session); err != nil {
      return nil, err
   }
   return session, nil
}

func (w *web_token) unmarshal() error {
   return json.Unmarshal(w.data, &w.v)
}
