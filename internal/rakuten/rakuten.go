package main

import (
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   fhd, err := f.address.Fhd().Info()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", fhd.Url, nil)
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
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name, err = f.address.Movie()
         if err != nil {
            return err
         }
         hd, err := f.address.Hd().Info()
         if err != nil {
            return err
         }
         fhd.LicenseUrl = hd.LicenseUrl
         f.s.Poster = fhd
         return f.s.Download(rep)
      }
   }
   return nil
}
