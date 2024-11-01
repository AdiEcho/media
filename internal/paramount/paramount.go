package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/paramount"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func (f *flags) do_read() error {
   data, err := os.ReadFile(f.content_id + "/request.txt")
   if err != nil {
      return err
   }
   var address url.URL
   err = address.UnmarshalBinary(data)
   if err != nil {
      return err
   }
   data, err = os.ReadFile(f.content_id + "/body.txt")
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, &address)
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
         data, err = os.ReadFile(f.content_id + "/item.txt")
         if err != nil {
            return err
         }
         var item paramount.VideoItem
         err = item.Unmarshal(data)
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
   os.Mkdir(f.content_id, os.ModePerm)
   // item
   var (
      app paramount.AppToken
      err error
   )
   if f.intl {
      err = app.ComCbsCa()
   } else {
      err = app.ComCbsApp()
   }
   if err != nil {
      return err
   }
   var data []byte
   _, err = app.Item(f.content_id, &data)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.content_id + "/item.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // mpd
   var item paramount.VideoItem
   err = item.Unmarshal(data)
   if err != nil {
      return err
   }
   resp, err := http.Get(item.Mpd())
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   // Body
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.content_id + "/body.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // Request
   data, err = resp.Request.URL.MarshalBinary()
   if err != nil {
      return err
   }
   return os.WriteFile(f.content_id + "/request.txt", data, os.ModePerm)
}
