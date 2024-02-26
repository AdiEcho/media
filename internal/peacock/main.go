package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/peacock"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   dash_id string
   email string
   h internal.HttpStream
   home string
   password string
   peacock_id string
   v log.Level
}

func main() {
   var (
      f flags
      home string
   )
   f.home, home = func() (string, string) {
      s, err := os.UserHomeDir()
      if err != nil {
         panic(err)
      }
      s = filepath.ToSlash(s)
      return s + "/", s + "/widevine/"
   }()
   flag.StringVar(&f.peacock_id, "b", "", "Peacock ID")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.dash_id, "i", "", "DASH ID")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.peacock_id != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
