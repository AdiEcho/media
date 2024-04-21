package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/pluto"
   "flag"
   "os"
   "path/filepath"
   "strings"
)

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.web, "a", "address")
   flag.StringVar(&f.base, "b", pluto.Base[0], func() string {
      var b strings.Builder
      for _, base := range pluto.Base[1:] {
         b.WriteString(base)
         b.WriteByte('\n')
      }
      return b.String()
   }())
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.forward, "f", "", func() string {
      var b strings.Builder
      for key, value := range pluto.Forward {
         if b.Len() >= 1 {
            b.WriteByte('\n')
         }
         b.WriteString(key)
         b.WriteByte(' ')
         b.WriteString(value)
      }
      return b.String()
   }())
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.web.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   forward string
   base string
   h internal.HttpStream
   representation string
   v log.Level
   web pluto.WebAddress
}

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
