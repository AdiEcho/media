package main

import "154.pages.dev/media/paramount"

func (f flags) dash(app paramount.AppToken) error {
   address, err := paramount.DashCenc(f.paramount_id)
   if err != nil {
      return err
   }
   media, err := f.h.DashMedia(address)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      f.h.Poster, err = app.Session(f.paramount_id)
      if err != nil {
         return err
      }
      item, err := app.Item(f.paramount_id)
      if err != nil {
         return err
      }
      f.h.Name = item
   }
   return f.h.DASH(media, f.dash_id)
}
