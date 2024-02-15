package main

import (
   "154.pages.dev/log"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   h rosso.HttpStream
   v log.Level
   guid int64
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.BoolVar(&f.h.Info, "i", false, "information")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.IntVar(&f.bandwidth, "m", 6_999_999, "max video bandwidth")
   flag.TextVar(&f.level, "v", f.level, "level")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
