package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

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

type flags struct {
   dash string
   h internal.HttpStream
   path plex.Path
   v log.Level
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.StringVar(&f.dash, "i", "", "representation ID")
   flag.Var(&f.path, "a", "plex path")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.path.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) download() error {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   match, err := anon.discover(tests["movie"])
   if err != nil {
      t.Fatal(err)
   }
   video, err := anon.on_demand(match)
   if err != nil {
      t.Fatal(err)
   }
   part, ok := video.dash(anon)
   if !ok {
      t.Fatal("metadata.dash")
   }
   // OLD
   var meta plex.Metadata
   err := meta.New(f.plex)
   if err != nil {
      return err
   }
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
      if medium.ID == f.dash {
         f.h.Name = meta
         f.h.Poster = plex.Core()
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
