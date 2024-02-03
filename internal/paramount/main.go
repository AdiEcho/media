package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/paramount"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int
   height int
   codec string
   role string
   content_ID string
   dash_cenc bool
   s stream.Stream
   level log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.StringVar(&f.codec, "ac", "mp4a", "audio codec")
   flag.StringVar(&f.role, "ar", "main", "audio role")
   flag.StringVar(&f.content_ID, "b", "", "content ID")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.BoolVar(&f.dash_cenc, "d", false, "DASH_CENC")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.IntVar(&f.bandwidth, "vb", 5_000_000, "video max bandwidth")
   flag.IntVar(&f.height, "vh", 720, "video max height")
   flag.TextVar(&f.level, "v", f.level, "level")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
   if f.content_ID != "" {
      token, err := paramount.New_App_Token()
      if err != nil {
         panic(err)
      }
      if f.dash_cenc {
         err = f.dash(token)
      } else {
         err = f.downloadable(token)
      }
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
