package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   roku string
   representation string
   s internal.Stream
   log text.LogLevel
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
   flag.StringVar(&f.roku, "b", "", "Roku ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.TextVar(&f.log.Level, "v", f.log.Level, "level")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
   if f.roku != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
