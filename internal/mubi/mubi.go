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

func TestFilm(t *testing.T) {
   for i, dogville := range dogvilles {
      var web WebAddress
      err := web.Set(dogville)
      if err != nil {
         t.Fatal(err)
      }
      if i == 0 {
         film, err := web.film()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(encoding.Name(film))
      }
      fmt.Println(web)
   }
}

func TestSecure(t *testing.T) {
   var (
      auth authenticate
      err error
   )
   auth.Raw, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   secure, err := auth.secure(passages_2022)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}

func TestAuthenticate(t *testing.T) {
   var (
      code link_code
      err error
   )
   code.Raw, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   auth, err := code.authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.Raw, 0666)
}

func TestCode(t *testing.T) {
   var code link_code
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Raw, 0666)
   code.unmarshal()
   fmt.Println(code)
}

func (f flags) download() error {
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
}
