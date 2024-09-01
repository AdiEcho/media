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
   address pluto.Address
   get_forward bool
   set_forward string
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
   flag.BoolVar(&f.get_forward, "g", false, "get forward")
   flag.StringVar(&f.set_forward, "s", "", "set forward")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.get_forward:
      get_forward()
   case f.address.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
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
