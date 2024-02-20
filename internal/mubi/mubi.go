package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/mubi"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   dash_id string
   h internal.HttpStream
   mubi_id int
   v log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.IntVar(&f.mubi_id, "b", 0, "Mubi ID")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.dash_id, "d", "", "DASH ID")
   flag.StringVar(&f.h.Private_Key, "p", home+"private_key.pem", "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   if f.mubi_id >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) download() error {
   var meta mubi.Metadata
   err := meta.New(f.mubi_id)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      f.h.Name = meta
      f.h.Poster = mubi.Core()
   }
   return f.h.DASH(media, f.dash_id)
}
