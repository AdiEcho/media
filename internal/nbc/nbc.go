package main

import "154.pages.dev/media/nbc"

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
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      f.h.Name = meta
      f.h.Poster = nbc.Core()
   }
   return f.h.DASH(media, f.dash_id)
}
