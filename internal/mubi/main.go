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
   dash_id string
   h internal.HttpStream
   home string
   v log.Level
   web mubi.WebAddress
   code bool
   auth bool
}

func main() {
   var (
      f flags
      home string
   )
   f.home, home = func() (string, string) {
      s, err := os.UserHomeDir()
      if err != nil {
         panic(err)
      }
      return s + "/mubi/", filepath.ToSlash(s) + "/widevine/"
   }()
   flag.Var(&f.web, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.h.Private_Key, "p", home+"private_key.pem", "private key")
   flag.StringVar(&f.dash_id, "d", "", "DASH ID")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
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
   case f.web.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
