package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
)

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.TextVar(&f.log.Level, "v", f.log.Level, "level")
   flag.StringVar(&f.forward, "z", "", internal.Forward.String())
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
   if f.address.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   address plex.Path
   representation string
   s internal.Stream
   log text.LogLevel
   forward string
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

