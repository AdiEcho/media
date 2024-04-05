package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/stan"
   "flag"
   "os"
   "path/filepath"
   "strings"
)

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.h.ClientId = f.home + "/widevine/client_id.bin"
   f.h.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   h internal.HttpStream
   home string
   representation string
   v log.Level
   program int64
   code bool
   token bool
   host string
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Int64Var(&f.program, "b", 0, "program ID")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.BoolVar(&f.code, "code", false, "activation code")
   flag.StringVar(
      &f.host, "h", stan.BaseUrl[0], strings.Join(stan.BaseUrl[1:], "\n"),
   )
   flag.StringVar(&f.representation, "i", "", "representation ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.BoolVar(&f.token, "token", false, "web token")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.code:
      err := f.write_code()
      if err != nil {
         panic(err)
      }
   case f.token:
      err := f.write_token()
      if err != nil {
         panic(err)
      }
   case f.program >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
