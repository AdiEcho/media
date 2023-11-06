package main

import (
   "154.pages.dev/media/nbc"
   "154.pages.dev/stream"
   "154.pages.dev/stream/dash"
   "net/http"
   "slices"
)

func (f flags) download() error {
   meta, err := nbc.New_Metadata(f.guid)
   if err != nil {
      return err
   }
   reps, err := func() ([]dash.Representation, error) {
      o, err := meta.On_Demand()
      if err != nil {
         return nil, err
      }
      r, err := http.Get(o.Playback_URL)
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
   slices.SortFunc(reps, func(a, b dash.Representation) int {
      return int(b.Bandwidth - a.Bandwidth)
   })
   index := slices.IndexFunc(reps, func(a dash.Representation) bool {
      return a.Bandwidth <= f.bandwidth
   })
   return f.s.DASH_Get(reps, index)
}
