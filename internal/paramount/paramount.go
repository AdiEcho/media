package main

import (
   "154.pages.dev/media/paramount"
   "fmt"
)

func (f flags) dash(app paramount.AppToken) error {
   address, err := paramount.DashCenc(f.paramount_id)
   if err != nil {
      return err
   }
   // 1 MPD one
   media, err := f.h.DashMedia(address)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         f.h.Name, err = app.Item(f.paramount_id)
         if err != nil {
            return err
         }
         f.h.Poster, err = app.Session(f.paramount_id)
         if err != nil {
            return err
         }
         return f.h.DASH(medium)
      }
   }
   // 2 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
