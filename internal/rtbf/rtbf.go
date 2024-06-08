package main

import (
   "154.pages.dev/media/rtbf"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
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
   
   token, err := account.token()
   if err != nil {
      t.Fatal(err)
   }
   gigya, err := token.login()
   if err != nil {
      t.Fatal(err)
   }
   page, err := new_page(media[0].path)
   if err != nil {
      t.Fatal(err)
   }
   title, err := gigya.entitlement(page)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", title)
   fmt.Println(title.dash())
   req, err := http.NewRequest("", play.StreamUrl, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name, err = auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Poster = play
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
   os.WriteFile(f.home + "/rtbf.json", text, 0666)
}
