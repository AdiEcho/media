package main

import (
   "154.pages.dev/log"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   guid int64
   bandwidth int
   h log.Handler
   s stream.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.IntVar(&f.bandwidth, "bandwidth", 6_999_999, "maximum bandwidth")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.TextVar(&f.h.Level, "v", f.h.Level, "level")
   flag.Parse()
   log.Set_Handler(f.h)
   log.Set_Transport(0)
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
