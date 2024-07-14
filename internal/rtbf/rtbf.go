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
   text, err := os.ReadFile(f.home + "/rtbf.json")
   if err != nil {
      return err
   }
   var account rtbf.AccountLogin
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
   media, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         f.s.Name = page
         f.s.Poster = title
         return f.s.Download(medium)
      }
   }
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}

func (f flags) authenticate() error {
   var login rtbf.AccountLogin
   err := login.New(f.email, f.password)
   if err != nil {
      return err
   }
   text, err := login.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/rtbf.json", text, 0666)
}
