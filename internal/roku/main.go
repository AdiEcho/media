package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/roku"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   roku_id string
   dash_id string
   h internal.HttpStream
   v log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.StringVar(&f.roku_id, "b", "", "Roku ID")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.dash_id, "i", "", "DASH ID")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   if f.roku_id != "" {
      var home roku.HomeScreen
      err := home.New(f.roku_id)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(home); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
