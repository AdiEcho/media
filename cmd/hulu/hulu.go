package main

import (
   "154.pages.dev/dash"
   "154.pages.dev/media/hulu"
   "154.pages.dev/stream"
   "net/http"
   "os"
   "slices"
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
   detail, err := auth.Details(deep)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   reps, err := func() ([]dash.Representation, error) {
      res, err := http.Get(play.Stream_URL)
      if err != nil {
         return nil, err
      }
      defer res.Body.Close()
      f.s.Base = res.Request.URL
      return dash.Representations(res.Body)
   }()
   if err != nil {
      return err
   }
   if !f.s.Info {
      f.s.Name, err = stream.Format_Episode(detail)
      if err != nil {
         return err
      }
      f.s.Poster = play
   }
   slices.SortFunc(reps, func(a, b dash.Representation) int {
      return b.Bandwidth - a.Bandwidth
   })
   // video
   {
      reps := slices.DeleteFunc(slices.Clone(reps), dash.Not(dash.Video))
      index := slices.IndexFunc(reps, func(a dash.Representation) bool {
         return a.Bandwidth <= f.bandwidth
      })
      err := f.s.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, dash.Not(dash.Audio))
   return f.s.DASH_Get(reps, 0)
}

func (f flags) authenticate() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/hulu/token.json"
   auth, err := hulu.Living_Room(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}
