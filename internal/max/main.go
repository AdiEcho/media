package main

import (
   "154.pages.dev/media/max"
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
   f.home += "/max"
   return nil
}

type flags struct {
   email string
   home string
   max_height int
   password string
   representation string
   s internal.Stream
   address max.Address
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.TextVar(&f.address, "a", &f.address, "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.IntVar(&f.max_height, "m", 1079, "max height")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.address.VideoId != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
