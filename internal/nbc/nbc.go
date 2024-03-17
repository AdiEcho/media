package main

import (
   "154.pages.dev/media/nbc"
   "fmt"
)

func (f flags) download() error {
   var meta nbc.Metadata
   err := meta.New(f.nbc_id)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   // 1 MPD one
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         f.h.Name = meta
         f.h.Poster = nbc.Core()
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
