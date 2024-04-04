package main

import (
   "154.pages.dev/media/stan"
   "fmt"
   "os"
)

func (f flags) write_code() error {
   var code stan.ActivationCode
   err := code.New()
   if err != nil {
      return err
   }
   code.Unmarshal()
   fmt.Println(code)
   os.WriteFile("code.json", code.Data, 0666)
}

func (f flags) write_token() error {
   var code stan.ActivationCode
   code.Data, err = os.ReadFile("code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   token, err := code.Token()
   if err != nil {
      return err
   }
   os.WriteFile(home + "/stan.json", token.Data, 0666)
}

func (f flags) download() error {
   var token stan.WebToken
   token.Data, err = os.ReadFile(f.home + "/stan.json")
   if err != nil {
      return err
   }
   
   token.unmarshal()
   session, err := token.Session()
   if err != nil {
      t.Fatal(err)
   }
   stream, err := session.Stream(program_id)
   if err != nil {
      t.Fatal(err)
   }
   // OLD
   var (
      secure stan.SecureUrl
      err error
   )
   secure.Data, err = os.ReadFile(f.web.String() + ".json")
   if err != nil {
      return err
   }
   secure.Unmarshal()
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
   // 4 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
