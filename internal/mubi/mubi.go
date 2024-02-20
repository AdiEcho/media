package main

import (
   "154.pages.dev/encoding"
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/mubi"
   "flag"
   "fmt"
   "os"
   "path/filepath"
   "testing"
)

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
   // FIXME
   secure, err := auth.Secure(passages_2022)
   if err != nil {
      t.Fatal(err)
   }
   // NEW END
   // OLD BEGIN
   var meta mubi.Metadata
   err := meta.New(f.mubi_id)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      f.h.Name = meta
      f.h.Poster = mubi.Core()
   }
   return f.h.DASH(media, f.dash_id)
   // OLD END
   // NEW BEGIN
   film, err := web.film()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(encoding.Name(film))
   // NEW END
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
   os.WriteFile(f.home + "authenticate.json", auth.Raw, 0666)
}
