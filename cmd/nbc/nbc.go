package main

import (
   "154.pages.dev/dash"
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
      return dash.Representations(r.Body)
   }()
   if err != nil {
      return err
   }
   if !f.s.Info {
      f.s.Name, err = stream.Format_Episode(meta)
      if err != nil {
         return err
      }
      f.s.Poster = nbc.Core
   }
   // video
   index := slices.IndexFunc(reps, func(r *dash.Representation) bool {
      return r.Bandwidth <= f.bandwidth
   })
   if err := f.s.DASH_Get(reps, index); err != nil {
      return err
   }
   // audio
   reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
      return !r.Audio()
   })
   return f.s.DASH_Get(reps, 0)
}
