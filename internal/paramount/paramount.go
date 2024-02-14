package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/media/paramount"
   "154.pages.dev/rosso"
   "fmt"
   "net/http"
   "os"
   "slices"
   "strings"
)

func (f flags) dash(app paramount.AppToken) error {
   ref, err := paramount.DashCenc(f.paramount_id)
   if err != nil {
      return err
   }
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f.s.Base = res.Request.URL
   if !f.s.Info {
      f.s.Poster, err = app.Session(f.paramount_id)
      if err != nil {
         return err
      }
      item, err := app.Item(f.paramount_id)
      if err != nil {
         return err
      }
      f.s.Name = stream.Name(item)
   }
   var media dash.Media
   media.Decode(res.Body)
   reps, err := media.Representation("0")
   if err != nil {
      return err
   }
   return f.s.DASH_Sofia(reps, index)
}

func (f flags) downloadable(app paramount.AppToken) error {
   item, err := app.Item(f.paramount_id)
   if err != nil {
      return err
   }
   ref, err := paramount.Downloadable(f.paramount_id)
   if err != nil {
      return err
   }
   if f.s.Info {
      fmt.Println(ref)
      return nil
   }
   dst, err := os.Create(stream.Name(item) + ".mp4")
   if err != nil {
      return err
   }
   defer dst.Close()
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   src := log.NewProgress(1).Reader(res)
   if _, err := dst.ReadFrom(src); err != nil {
      return err
   }
   return nil
}
