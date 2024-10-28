package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/paramount"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f *flags) do_write() error {
   os.Mkdir(f.content_id, os.ModePerm)
   location := locations[f.location]
   mpd, err := paramount.Mpd(f.content_id, location.asset_type)
   if err != nil {
      return err
   }
   err = os.WriteFile(
      f.content_id + "/mpd.txt", []byte(mpd), os.ModePerm,
   )
   if err != nil {
      return err
   }
   var app paramount.AppToken
   if location.host == "www.paramountplus.com" {
      err = app.ComCbsApp()
   } else {
      err = app.ComCbsCa()
   }
   if err != nil {
      return err
   }
   item, err := app.Item(f.content_id)
   if err != nil {
      return err
   }
   return os.WriteFile(f.content_id + "/item.txt", item.Raw, os.ModePerm)
}

func (f *flags) do_read() error {
   mpd, err := os.ReadFile(f.content_id + "/mpd.txt")
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", string(mpd), nil)
   if err != nil {
      return err
   }
   reps, err := internal.Mpd(req)
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
         f.s.Poster, err = app.Session(f.content_id)
         if err != nil {
            return err
         }
         var item paramount.VideoItem
         item.Raw, err = os.ReadFile(f.content_id + "/item.txt")
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
