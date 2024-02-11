package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/youtube"
   "flag"
   "strings"
)

type flags struct {
   refresh bool
   r youtube.Request
   request int
   v log.Level
}

func main() {
   var f flags
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.r.Video_ID, "b", "", "video ID")
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   switch {
   case f.r.Video_ID != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   case f.refresh:
      err := f.do_refresh()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
