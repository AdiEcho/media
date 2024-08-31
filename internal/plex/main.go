package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   forward string
   representation string
   s internal.Stream
   url plex.Url
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.StringVar(&f.forward, "z", "", internal.Forward.String())
   flag.Var(&f.url, "a", "address")
   flag.Parse()
   text.Transport{}.Set(true)
   if f.url.Path != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
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

