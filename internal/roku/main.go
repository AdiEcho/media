package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/roku"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   roku_id string
   dash_id string
   h rosso.HttpStream
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
   flag.StringVar(&f.dash_id, "d", "", "DASH ID")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   if f.roku_id != "" {
      content, err := roku.New_Content(f.id)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(content); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
