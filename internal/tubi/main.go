package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   representation string
   s internal.Stream
   tubi int
   v log.Level
   content bool
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
   flag.IntVar(&f.tubi, "b", 0, "Tubi ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.BoolVar(&f.content, "w", false, "write content")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.content:
      err := f.write_content()
      if err != nil {
         panic(err)
      }
   case f.tubi >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
