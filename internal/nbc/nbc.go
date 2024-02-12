package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/nbc"
   "154.pages.dev/rosso"
   "net/http"
   "slices"
)

func (f flags) download() error {
   meta, err := nbc.NewMetadata(f.guid)
   if err != nil {
      return err
   }
   reps, err := func() ([]*dash.Representation, error) {
      on, err := meta.OnDemand()
      if err != nil {
         return nil, err
      }
      r, err := http.Get(on.PlaybackUrl)
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
      f.s.Name = rosso.Name(meta)
   }
   return f.s.DASH_Sofia(reps, index)
}
