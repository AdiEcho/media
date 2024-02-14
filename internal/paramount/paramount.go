package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/paramount"
   "154.pages.dev/rosso"
   "net/http"
)

func (f flags) dash(app paramount.AppToken) error {
   address, err := paramount.DashCenc(f.paramount_id)
   if err != nil {
      return err
   }
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f.h.Base = res.Request.URL
   if !f.h.Info {
      f.h.Poster, err = app.Session(f.paramount_id)
      if err != nil {
         return err
      }
      item, err := app.Item(f.paramount_id)
      if err != nil {
         return err
      }
      f.h.Name = rosso.Name(item)
   }
   var media dash.MPD
   media.Decode(res.Body)
   reps, err := media.Representation("0")
   if err != nil {
      return err
   }
   return f.h.DASH_Sofia(reps, index)
}
