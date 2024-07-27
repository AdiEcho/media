package main

import (
   "154.pages.dev/media/cine/member"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
)

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.s.ClientId = f.home + "/widevine/client_id.bin"
   f.s.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   email string
   s internal.Stream
   home string
   representation string
   password string
   slug member.ArticleSlug
   play bool
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.slug, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.play, "o", false, "operation play")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.password != "":
      err := f.write_user()
      if err != nil {
         panic(err)
      }
   case f.play:
      err := f.write_play()
      if err != nil {
         panic(err)
      }
   case f.slug != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
