package main

import (
   "41.neocities.org/dash"
   "fmt"
   "io"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   fhd, err := f.address.Fhd().Info()
   if err != nil {
      return err
   }
   resp, err := http.Get(fhd.Url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
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
         f.s.Wrapper = fhd
         return f.s.Download(rep)
      }
   }
   return nil
}
