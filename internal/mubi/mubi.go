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
   code.Data, err = os.ReadFile(f.home + "link_code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "authenticate.json", auth.Data, 0666)
}

func (f flags) write_code() error {
   var code mubi.LinkCode
   err := code.New()
   if err != nil {
      return err
   }
   os.WriteFile(f.home + "link_code.json", code.Data, 0666)
   code.Unmarshal()
   fmt.Println(code)
   return nil
}

func (f flags) write_secure() error {
   var (
      auth mubi.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "authenticate.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   film, err := f.web.Film()
   if err != nil {
      return err
   }
   if err := auth.Viewing(film); err != nil {
      return err
   }
   secure, err := auth.URL(film)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + f.web.String(), secure.Data, 0666)
}

func (f flags) download() error {
   film, err := f.web.Film()
   if err != nil {
      return err
   }
   var auth mubi.Authenticate
   auth.Data, err = os.ReadFile(f.home + "authenticate.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   var secure mubi.SecureUrl
   secure.Data, err = os.ReadFile(f.home + f.web.String())
   if err != nil {
      return err
   }
   secure.Unmarshal()
   media, err := f.h.DashMedia(secure.V.URL)
   if err != nil {
      return err
   }
   f.h.Name = film
   f.h.Poster = auth
   return f.h.DASH(media, f.dash_id)
}
