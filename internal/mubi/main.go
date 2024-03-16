package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/mubi"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   media_id string
   h internal.HttpStream
   v log.Level
   web mubi.WebAddress
   code bool
   auth bool
   secure bool
   home string
}

func main() {
   var (
      f flags
      err error
   )
   f.home, err = os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   f.home = filepath.ToSlash(f.home)
   home := f.home + "/widevine"
   flag.Var(&f.web, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.StringVar(&f.h.Client_ID, "c", home+"/client_id.bin", "client ID")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.StringVar(&f.h.Private_Key, "p", home+"/private_key.pem", "private key")
   flag.BoolVar(&f.secure, "s", false, "secure URL")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
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
