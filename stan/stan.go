package stan

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type ActivationCode struct {
   Data []byte
   V struct {
      Code string
      URL string
   }
}

func (a *ActivationCode) New() error {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/activation-codes/", url.Values{
         "generate": {"true"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a ActivationCode) String() string {
   var b strings.Builder
   b.WriteString("Stan.\n")
   b.WriteString("Log in with code\n")
   b.WriteString("1. Visit stan.com.au/activate\n")
   b.WriteString("2. Enter the code:\n")
   b.WriteString(a.V.Code)
   return b.String()
}

func (a ActivationCode) token() (*web_token, error) {
   res, err := http.Get(a.V.URL)
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
   web.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &web, nil
}

func (a *ActivationCode) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

type app_session struct {
   JwToken string
}

type web_token struct {
   Data []byte
   V struct {
      JwToken string
      ProfileId string
   }
}

func (w web_token) session() (*app_session, error) {
   res, err := http.PostForm(
      "https://api.stan.com.au/login/v1/sessions/mobile/app", url.Values{
         "jwToken": {w.V.JwToken},
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
   return json.Unmarshal(w.Data, &w.V)
}
