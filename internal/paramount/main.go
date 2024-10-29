package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/text"
   "flag"
   "net/url"
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
   representation string
   s internal.Stream
   content_id string
   write bool
   intl bool
   url *url.URL
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.content_id, "b", "", "content ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.write, "w", false, "write")
   flag.BoolVar(&f.intl, "n", false, "intl")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.write:
      err := f.do_write()
      if err != nil {
         panic(err)
      }
   case f.content_id != "":
      err := f.do_read()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
