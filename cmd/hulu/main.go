package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/hulu"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   email string
   h log.Handler
   id hulu.ID
   password string
   s stream.Stream
   video_bandwidth int
   audio_codec string
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.id, "a", "address")
   flag.StringVar(&f.audio_codec, "ac", "ec-3", "audio codec")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.h.Level, "v", f.h.Level, "level")
   flag.IntVar(&f.video_bandwidth, "vb", 8_500_000, "video max bandwidth")
   flag.Parse()
   log.Set_Handler(f.h)
   log.Set_Transport(0)
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
