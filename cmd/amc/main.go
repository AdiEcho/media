package main

import (
   "154.pages.dev/media/amc"
   "154.pages.dev/log"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   email string
   height int
   password string
   address amc.URL
   s stream.Stream
   h log.Handler
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.h.Level, "v", f.h.Level, "log level")
   flag.Parse()
   log.Set_Handler(f.h)
   log.Set_Transport(0)
   switch {
   case f.email != "":
      err := f.login()
      if err != nil {
         panic(err)
      }
   case f.address.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
