package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/pluto"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
   "strings"
)

type flags struct {
   base string
   s internal.Stream
   representation string
   log text.LogLevel
   address pluto.Address
   forward string
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.base, "b", pluto.Base[0], func() string {
      var b strings.Builder
      for _, base := range pluto.Base[1:] {
         b.WriteString(base)
         b.WriteByte('\n')
      }
      return b.String()
   }())
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
