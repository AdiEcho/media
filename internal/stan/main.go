package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/stan"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   media_id string
   h internal.HttpStream
   v log.Level
   web stan.WebAddress
   code bool
   auth bool
   secure bool
   home string
}

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

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.web, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.BoolVar(&f.secure, "s", false, "secure URL")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.auth:
      err := f.write_auth()
      if err != nil {
         panic(err)
      }
   case f.code:
      err := f.write_code()
      if err != nil {
         panic(err)
      }
   case f.secure:
      err := f.write_secure()
      if err != nil {
         panic(err)
      }
   case f.web.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
