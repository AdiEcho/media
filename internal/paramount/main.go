package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/paramount"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   representation string
   s internal.Stream
   paramount string
   v text.Level
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.paramount, "b", "", "Paramount ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   f.v.Set()
   text.Transport{}.Set()
   if f.paramount != "" {
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
