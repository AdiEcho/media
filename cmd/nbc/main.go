package main

import (
   "154.pages.dev/http/option"
   "154.pages.dev/stream"
   "flag"
)

type flags struct {
   guid int64
   bandwidth int64
   resolution string
   s stream.Stream
   trace bool
}

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.Int64Var(&f.bandwidth, "bandwidth", 8_299_999, "maximum bandwidth")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.resolution, "r", "720", "resolution")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   option.No_Location()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
   }
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
