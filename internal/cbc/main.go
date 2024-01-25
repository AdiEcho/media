package main

import (
   "154.pages.dev/log"
   "154.pages.dev/stream"
   "flag"
)

type flags struct {
   address string
   audio_name string
   email string
   password string
   s stream.Stream
   video_height string
   video_rate int64
   level log.Level
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Int64Var(&f.video_rate, "b", 3_000_000, "max video bandwidth")
   flag.StringVar(&f.video_height, "r", "720", "video resolution")
   flag.StringVar(&f.audio_name, "n", "English", "audio name")
   flag.TextVar(&f.level, "v", f.level, "log level")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
   if f.email != "" {
      err := f.profile()
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
