package main

import (
   "154.pages.dev/media/stan"
   "fmt"
   "os"
)

func (f flags) write_code() error {
   var code activation_code
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   fmt.Println(code)
   os.WriteFile("code.json", code.data, 0666)
}

func (f flags) write_auth() error {
   var (
      code stan.LinkCode
      err error
   )
   code.Data, err = os.ReadFile("link_code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/stan.json", auth.Data, 0666)
}

func (f flags) download() error {
   var (
      secure stan.SecureUrl
      err error
   )
   secure.Data, err = os.ReadFile(f.web.String() + ".json")
   if err != nil {
      return err
   }
   secure.Unmarshal()
   // 1 VTT one
   for _, text := range secure.V.Text_Track_URLs {
      if text.ID == f.media_id {
         f.h.Name, err = f.web.Film()
         if err != nil {
            return err
         }
         return f.h.TimedText(text.URL)
      }
   }
   // 2 MPD one
   media, err := f.h.DashMedia(secure.V.URL)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         f.h.Name, err = f.web.Film()
         if err != nil {
            return err
         }
         var auth stan.Authenticate
         auth.Data, err = os.ReadFile(f.home + "/stan.json")
         if err != nil {
            return err
         }
         auth.Unmarshal()
         f.h.Poster = auth
         return f.h.DASH(medium)
      }
   }
   // 3 VTT all
   for _, text := range secure.V.Text_Track_URLs {
      fmt.Print(text, "\n\n")
   }
   // 4 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
