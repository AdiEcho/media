package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/hulu"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   media_id string
   email string
   h internal.HttpStream
   hulu_id hulu.ID
   password string
   v log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine"
   var f flags
   flag.Var(&f.hulu_id, "a", "address")
   flag.StringVar(&f.h.Client_ID, "c", home+"/client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.StringVar(&f.h.Private_Key, "k", home+"/private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.hulu_id.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
