package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/nbc"
   "154.pages.dev/stream"
   "net/http"
   "slices"
)

func (f flags) download() error {
   meta, err := nbc.New_Metadata(f.guid)
   if err != nil {
      return err
   }
   reps, err := func() ([]*dash.Representation, error) {
      on, err := meta.On_Demand()
      if err != nil {
         return nil, err
      }
      r, err := http.Get(on.Playback_URL)
      if err != nil {
         return nil, err
      }
      defer r.Body.Close()
      f.s.Base = r.Request.URL
      var media dash.Media
      media.Decode(r.Body)
      return media.Representation("")
   }()
   if err != nil {
      return err
   }
   if !f.s.Info {
      f.s.Poster = nbc.Core
      f.s.Name = stream.Name(meta)
   }
   slices.SortFunc(reps, func(a, b *dash.Representation) int {
      return b.Bandwidth - a.Bandwidth
   })
   // video
   {
      reps := slices.Clone(reps)
      reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
         return !r.Video()
      })
      index := slices.IndexFunc(reps, func(r *dash.Representation) bool {
         return r.Bandwidth <= f.bandwidth
      })
      err := f.s.DASH_Sofia(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
      return !r.Audio()
   })
   return f.s.DASH_Sofia(reps, 1)
}
