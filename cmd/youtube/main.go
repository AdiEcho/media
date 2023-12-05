package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/youtube"
   "flag"
   "strings"
)

type flags struct {
   a_codec string
   a_quality string
   info bool
   refresh bool
   r youtube.Request
   request int
   v_codec string
   v_quality string
   h log.Handler
}

func main() {
   var f flags
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.a_codec, "ac", "opus", "audio codec")
   flag.StringVar(&f.a_quality, "aq", "AUDIO_QUALITY_MEDIUM", "audio quality")
   flag.StringVar(&f.r.Video_ID, "b", "", "video ID")
   flag.BoolVar(&f.info, "i", false, "information")
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.TextVar(&f.h.Level, "v", f.h.Level, "level")
   flag.StringVar(&f.v_codec, "vc", "vp9", "video codec")
   flag.StringVar(&f.v_quality, "vq", "1080p", "video quality")
   flag.Parse()
   log.Set_Handler(f.h)
   log.Set_Transport(0)
   if f.refresh {
      err := f.do_refresh()
      if err != nil {
         panic(err)
      }
   } else if f.r.Video_ID != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
