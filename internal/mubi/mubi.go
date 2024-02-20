package main

import (
   "154.pages.dev/media/mubi"
   "fmt"
   "os"
)

func (f flags) write_auth() error {
   var (
      code mubi.LinkCode
      err error
   )
   code.Raw, err = os.ReadFile(f.home + "link_code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "authenticate.json", auth.Raw, 0666)
}

func (f flags) download() error {
   var (
      auth mubi.Authenticate
      err error
   )
   auth.Raw, err = os.ReadFile(f.home + "authenticate.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   film, err := f.web.Film()
   if err != nil {
      return err
   }
   secure, err := auth.URL(film)
   if err != nil {
      return err
   }
   media, err := f.h.DashMedia(secure.URL)
   if err != nil {
      return err
   }
   f.h.Name = film
   f.h.Poster = auth
   return f.h.DASH(media, f.dash_id)
}

func (f flags) write_code() error {
   var code mubi.LinkCode
   err := code.New()
   if err != nil {
      return err
   }
   os.WriteFile(f.home + "link_code.json", code.Raw, 0666)
   code.Unmarshal()
   fmt.Println(code)
   return nil
}
