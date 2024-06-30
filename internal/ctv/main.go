package main

import (
   "154.pages.dev/media/ctv"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
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
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   path ctv.Path
   representation string
   s internal.Stream
   manifest bool
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.path, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.BoolVar(&f.manifest, "m", false, "manifest")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.manifest:
      err := f.get_manifest()
      if err != nil {
         panic(err)
      }
   case f.path != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
