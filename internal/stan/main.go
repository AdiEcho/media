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

type flags struct {
   code bool
   home string
   host string
   stan int64
   representation string
   s internal.Stream
   token bool
   v log.Level
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.s.ClientId = f.home + "/widevine/client_id.bin"
   f.s.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Int64Var(&f.stan, "b", 0, "Stan ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.BoolVar(&f.code, "code", false, "activation code")
   flag.StringVar(
      &f.host, "h", stan.BaseUrl[0], strings.Join(stan.BaseUrl[1:], "\n"),
   )
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
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
   case f.stan >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
