package main

import (
   "154.pages.dev/media/rtbf"
   "154.pages.dev/media/internal"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

func (f flags) download() error {
   var account rtbf.LoginToken
   text, err := os.ReadFile(f.home + "/rtbf.txt")
   if err != nil {
      return err
   }
   err = account.Unmarshal(text)
   if err != nil {
      return err
   }
   token, err := account.Token()
   if err != nil {
      return err
   }
   gigya, err := token.Login()
   if err != nil {
      return err
   }
   address, err := url.Parse(f.address)
   if err != nil {
      return err
   }
   page, err := rtbf.NewPage(address.Path)
   if err != nil {
      return err
   }
   title, err := gigya.Entitlement(page)
   if err != nil {
      return err
   }
   locator, ok := title.Dash()
   if !ok {
      return errors.New("Entitlement.Dash")
   }
   req, err := http.NewRequest("", locator, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = page
         f.s.Poster = title
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f flags) authenticate() error {
   var login rtbf.LoginToken
   err := login.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/rtbf.txt", login.Raw, 0666)
}
