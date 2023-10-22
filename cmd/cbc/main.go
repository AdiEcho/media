package main

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "flag"
)

type flags struct {
   address string
   audio_name string
   email string
   password string
   s stream.Stream
   trace bool
   video_height string
   video_rate int64
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.Int64Var(&f.video_rate, "b", 3_000_000, "max video bandwidth")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.audio_name, "n", "English", "audio name")
   flag.StringVar(&f.password, "p", "", "password")
   flag.StringVar(&f.video_height, "r", "720", "video resolution")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   http.No_Location()
   if f.trace {
      http.Trace()
   } else {
      http.Verbose()
   }
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
