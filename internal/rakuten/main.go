package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/rakuten"
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
   s internal.Stream
   representation string
   log text.LogLevel
   address rakuten.WebAddress
   streamings bool
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.streamings, "s", false, "streamings")
   flag.TextVar(&f.log.Level, "v", f.log.Level, "level")
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
   switch {
   case f.streamings:
      err := f.write_stream()
      if err != nil {
         panic(err)
      }
   case f.address.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
