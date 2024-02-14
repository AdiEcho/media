package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/paramount"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   dash_cenc bool
   dash_id string
   h rosso.HttpStream
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
   flag.BoolVar(&f.dash_cenc, "d", false, "DASH_CENC")
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
      if f.dash_cenc {
         err = f.dash(app)
      } else {
         err = f.downloadable(app)
      }
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
