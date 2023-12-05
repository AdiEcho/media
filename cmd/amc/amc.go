package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/amc"
   "154.pages.dev/stream"
   "errors"
   "net/http"
   "os"
   "slices"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth amc.Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         return err
      }
      auth.Unmarshal(b)
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         return err
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
   }
   if !f.s.Info {
      content, err := auth.Content(f.path)
      if err != nil {
         return err
      }
      video, err := content.Video()
      if err != nil {
         return err
      }
      f.s.Name, err = stream.Format_Film(video)
      if err != nil {
         return err
      }
   }
   play, err := auth.Playback(f.path)
   if err != nil {
      return err
   }
   f.s.Poster = play
   res, err := http.Get(play.HTTP_DASH().Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   reps, err := dash.Representations(res.Body)
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
         return r.Height <= f.height
      })
      if err := f.s.DASH_Get(reps, index); err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
      return !r.Audio()
   })
   return f.s.DASH_Get(reps, 0)
}

func (f flags) login() error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(f.email, f.password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         return err
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
   }
   return nil
}

