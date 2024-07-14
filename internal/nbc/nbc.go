package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/nbc"
   "fmt"
   "net/http"
)

func (f flags) download() error {
   var meta nbc.Metadata
   err := meta.New(f.nbc)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", demand.PlaybackUrl, nil)
   if err != nil {
      return err
   }
   media, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         f.s.Name = meta
         f.s.Poster = nbc.Core()
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
