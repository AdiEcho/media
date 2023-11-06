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
   resolution string
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
   flag.IntVar(&f.bandwidth, "bandwidth", 6_999_999, "maximum bandwidth")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Client_ID, "client", home+"client_id.bin", "client ID")
   flag.StringVar(&f.s.Private_Key, "key", home+"private_key.pem", "private key")
   flag.Parse()
   http.No_Location()
   http.Verbose()
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
