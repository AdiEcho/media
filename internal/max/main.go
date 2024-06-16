package main

import (
   "154.pages.dev/media/max"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

type flags struct {
   email string
   s internal.Stream
   home string
   representation string
   password string
   log text.LogLevel
   address max.WebAddress
}

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

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.log.Level, "v", f.log.Level, "level")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.TextVar(f.address, "a", f.address, "address")
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
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
