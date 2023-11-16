package main

import (
   "154.pages.dev/http"
   "154.pages.dev/media/hulu"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int
   email string
   id hulu.ID
   password string
   s stream.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.id, "a", "address")
   flag.IntVar(&f.bandwidth, "b", 6_999_999, "maximum bandwidth")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   http.No_Location()
   http.Verbose()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.id.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
