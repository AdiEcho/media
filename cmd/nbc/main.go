package main

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   guid int64
   bandwidth int
   s stream.Stream
   trace bool
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
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   http.No_Location()
   if f.trace {
      http.Trace()
   } else {
      http.Verbose()
   }
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
