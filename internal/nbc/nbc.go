package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/nbc"
   "fmt"
   "net/http"
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
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = meta
         f.s.Poster = nbc.Core()
         return f.s.Download(rep)
      }
   }
   return nil
}
