package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/nbc"
   "slices"
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
   media, err := f.h.DashMedia(demand.PlaybackUrl)
   if err != nil {
      return err
   }
   f.h.Name = meta
   f.h.Poster = nbc.Core()
   for _, p := range media.Period {
      for _, a := range p.AdaptationSet {
         slices.SortFunc(a.Representation, func(a, b dash.Representation) int {
            return b.Bandwidth - a.Bandwidth
         })
      }
   }
   return f.h.DASH(media, f.dash_id)
}
