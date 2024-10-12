package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/paramount"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f *flags) do_read() error {
   location, err := os.ReadFile(f.paramount + "/location.txt")
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", string(location), nil)
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
         var app paramount.AppToken
         // INTL does NOT allow anonymous key request, so if you are INTL you
         // will need to use US VPN until someone codes the INTL login
         err := app.ComCbsApp()
         if err != nil {
            return err
         }
         f.s.Poster, err = app.Session(f.paramount)
         if err != nil {
            return err
         }
         var item paramount.VideoItem
         item.Raw, err = os.ReadFile(f.paramount + "/item.txt")
         if err != nil {
            return err
         }
         err = item.Unmarshal()
         if err != nil {
            return err
         }
         f.s.Name = &item
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) do_write() error {
   os.Mkdir(f.paramount, os.ModePerm)
   location, err := paramount.Location(f.paramount, f.intl)
   if err != nil {
      return err
   }
   err = os.WriteFile(
      f.paramount + "/location.txt", []byte(location), os.ModePerm,
   )
   if err != nil {
      return err
   }
   var app paramount.AppToken
   if f.intl {
      err = app.ComCbsCa()
   } else {
      err = app.ComCbsApp()
   }
   if err != nil {
      return err
   }
   item, err := app.Item(f.paramount)
   if err != nil {
      return err
   }
   return os.WriteFile(f.paramount + "/item.txt", item.Raw, os.ModePerm)
}
