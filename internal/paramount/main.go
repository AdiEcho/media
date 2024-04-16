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
   media_id string
   h internal.HttpStream
   paramount_id string
   v log.Level
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.h.ClientId = home + "/widevine/client_id.bin"
   f.h.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.paramount_id, "b", "", "Paramount ID")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.paramount_id != "" {
      var app paramount.AppToken
      err := app.New()
      if err != nil {
         panic(err)
      }
      err = f.dash(app)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
