package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/paramount"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   dash_id string
   h internal.HttpStream
   paramount_id string
   v log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.StringVar(&f.paramount_id, "b", "", "Paramount ID")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.dash_id, "i", "", "DASH ID")
   flag.StringVar(&f.h.Private_Key, "p", home+"private_key.pem", "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   if f.paramount_id != "" {
      var app paramount.AppToken
      err := app.New()
      if err != nil {
         panic(err)
      }
      if err := f.dash(app); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
