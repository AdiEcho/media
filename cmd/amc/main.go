package main

import (
   "154.pages.dev/http/option"
   "154.pages.dev/media"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   address string
   email string
   height int
   password string
   s media.Stream
}

func main() {
   home, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      return filepath.ToSlash(s) + "/widevine/", nil
   }()
   if err != nil {
      panic(err)
   }
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.s.Client_ID, "client", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "key", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   option.No_Location()
   option.Verbose()
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
