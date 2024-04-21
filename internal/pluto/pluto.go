package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/pluto"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

type flags struct {
   representation string
   h internal.HttpStream
   v log.Level
   web pluto.WebAddress
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.h.ClientId = home + "/widevine/client_id.bin"
   f.h.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.web, "a", "address")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.address.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) download() error {
   video, err := f.web.Video()
   if err != nil {
      return err
   }
   clip, err := video.Clip()
   if err != nil {
      return err
   }
   // OLD
   var meta pluto.Metadata
   err := meta.New(f.address)
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   // 1 MPD one
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         f.h.Name = meta
         f.h.Poster = pluto.Core()
         return f.h.DASH(medium)
      }
   }
   // 2 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
