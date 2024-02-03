package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/hulu"
   "154.pages.dev/rosso"
   "net/http"
   "os"
   "slices"
   "strings"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth hulu.Authenticate
   auth.Raw, err = os.ReadFile(home + "/hulu/token.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.Deep_Link(f.id)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   if !f.s.Info {
      detail, err := auth.Details(deep)
      if err != nil {
         return err
      }
      f.s.Name = rosso.Name(detail)
      f.s.Poster = play
   }
   
   
   
   
   
   
   reps, err := func() ([]*dash.Representation, error) {
      res, err := http.Get(play.Stream_URL)
      if err != nil {
         return nil, err
      }
      defer res.Body.Close()
      f.s.Base = res.Request.URL
      var media dash.Media
      media.Decode(res.Body)
      return media.Representation("content-0")
   }()
   if err != nil {
      return err
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
         return r.Bandwidth <= f.video_bandwidth
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
   index := slices.IndexFunc(reps, func(r *dash.Representation) bool {
      return strings.HasPrefix(r.Codecs, f.audio_codec)
   })
   return f.s.DASH_Sofia(reps, index)
}
