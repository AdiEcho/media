package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/text"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   representation string
   s internal.Stream
   content_id string
   write bool
   location int
}

var locations = []struct{
   host string
   asset_type string
}{
   {"www.paramountplus.com", "DASH_CENC"},
   {"www.intl.paramountplus.com", "DASH_CENC"},
   {"www.intl.paramountplus.com", "DASH_CENC_PRECON"},
}

func get_location() string {
   var b []byte
   for k, v := range locations {
      if k >= 1 {
         b = append(b, '\n')
      }
      b = fmt.Append(b, k, v)
   }
   return string(b)
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.content_id, "b", "", "content ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.IntVar(&f.location, "n", 0, get_location())
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.write, "w", false, "write")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.write:
      err := f.do_write()
      if err != nil {
         panic(err)
      }
   case f.content_id != "":
      err := f.do_read()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
