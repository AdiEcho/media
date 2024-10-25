package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/itv"
   "41.neocities.org/text"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
   "sort"
)

type flags struct {
   itv int
   representation string
   s internal.Stream
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.IntVar(&f.itv, "b", 0, "itv ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   text.Transport{}.Set(true)
   if f.itv >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
func (f *flags) download() error {
   var meta itv.Metadata
   err := meta.New(f.itv)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", demand.PlaybackUrl, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         if _, ok := rep.Ext(); ok {
            fmt.Print(&rep, "\n\n")
         }
      case rep.Id:
         f.s.Name = &meta
         var core itv.CoreVideo
         core.New()
         f.s.Poster = &core
         return f.s.Download(rep)
      }
   }
   return nil
}
