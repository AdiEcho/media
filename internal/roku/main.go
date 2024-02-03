package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/roku"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   s stream.Stream
   id string
   bandwidth int
   codec string
   height int
   lang string
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
   flag.StringVar(&f.lang, "al", "en", "audio language")
   flag.StringVar(&f.id, "b", "", "ID")
   flag.StringVar(&f.s.Client_ID, "c", home + "client_id.bin", "client ID")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.IntVar(&f.bandwidth, "vb", 4_200_000, "video max bandwidth")
   flag.IntVar(&f.height, "vh", 1080, "video max height")
   flag.TextVar(&f.level, "v", f.level, "level")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
   if f.id != "" {
      content, err := roku.New_Content(f.id)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(content); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
