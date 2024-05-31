package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/paramount"
   "fmt"
   "net/http"
)

func (f flags) dash(app paramount.AppToken) error {
   address, err := paramount.DashCenc(f.paramount)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Poster, err = app.Session(f.paramount)
         if err != nil {
            return err
         }
         item, err := app.Item(f.paramount)
         if err != nil {
            return err
         }
         f.s.Name = <-item
         return f.s.Download(medium)
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
