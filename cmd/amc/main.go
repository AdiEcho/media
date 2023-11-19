package main

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   address string
   email string
   height int
   password string
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
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   http.No_Location()
   if f.trace {
      http.Trace()
   } else {
      http.Verbose()
   }
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.address != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
