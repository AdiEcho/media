package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/paramount"
   "fmt"
   "net/http"
   "sort"
)

func (f flags) download() error {
   var app paramount.AppToken
   err := app.ComCbsApp()
   if err != nil {
      return err
   }
   // GEO
   address, err := paramount.MpegDash(f.paramount)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return err
   }
   reps, err := internal.DASH(req)
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
         // GEO
         f.s.Name, err = app.Item(f.paramount)
         if err != nil {
            return err
         }
         f.s.Poster, err = app.Session(f.paramount)
         if err != nil {
            return err
         }
         return f.s.Download(rep)
      }
   }
   return nil
}
