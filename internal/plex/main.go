package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "flag"
   "os"
   "path/filepath"
)

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

type flags struct {
   dash string
   h internal.HttpStream
   path plex.Path
   v log.Level
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.path, "a", "plex path")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&plex.Forward, "f", "", internal.Forward.String())
   flag.StringVar(&f.dash, "i", "", "representation ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.path.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
