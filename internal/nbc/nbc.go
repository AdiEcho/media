package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "fmt"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   var meta nbc.Metadata
   err := meta.New(f.nbc)
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
            fmt.Print(rep, "\n\n")
         }
      case rep.Id:
         f.s.Name = &meta
         var core nbc.CoreVideo
         core.New()
         f.s.Poster = &core
         return f.s.Download(rep)
      }
   }
   return nil
}
