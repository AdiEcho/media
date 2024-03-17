package main

import (
   "154.pages.dev/media/amc"
   "154.pages.dev/media/internal"
   "154.pages.dev/log"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   media_id string
   email string
   h internal.HttpStream
   password string
   v log.Level
   web amc.WebAddress
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.web, "a", "address")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "log level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.email != "":
      err := f.login()
      if err != nil {
         panic(err)
      }
   case f.web.NID != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
