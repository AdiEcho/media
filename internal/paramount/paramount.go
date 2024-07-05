package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/paramount"
   "errors"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f flags) do_read() error {
   text, err := os.ReadFile(f.paramount + "/header.json")
   if err != nil {
      return err
   }
   var head paramount.Header
   err = head.Json(text)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", head.Location(), nil)
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
         var app paramount.AppToken
         err := app.ComCbsApp()
         if err != nil {
            return err
         }
         f.s.Poster, err = app.Session(f.paramount)
         if err != nil {
            return err
         }
         text, err := os.ReadFile(f.paramount + "/item.json")
         if err != nil {
            return err
         }
         var item paramount.VideoItem
         err = item.Json(text)
         if err != nil {
            return err
         }
         f.s.Name = item
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f flags) do_write() error {
   err := os.Mkdir(f.paramount, 0666)
   if err != nil {
      return err
   }
   var head paramount.Header
   err = head.New(f.paramount)
   if err != nil {
      return err
   }
   text, err := head.JsonMarshal()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.paramount + "/header.json", text, 0666)
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
   items, err := app.Items(f.paramount)
   if err != nil {
      return err
   }
   item, ok := items.Item()
   if !ok {
      return errors.New("VideoItems.Item")
   }
   text, err = item.JsonMarshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.paramount + "/item.json", text, 0666)
}
