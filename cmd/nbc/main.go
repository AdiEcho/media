package main

import (
   "154.pages.dev/http/option"
   "154.pages.dev/media"
   "flag"
)

type flags struct {
   bandwidth int64
   guid int64
   resolution string
   s media.Stream
}

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.resolution, "r", "720", "resolution")
   flag.Int64Var(&f.bandwidth, "bandwidth", 8_099_999, "maximum bandwidth")
   flag.Parse()
   option.No_Location()
   option.Verbose()
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
