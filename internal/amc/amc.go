package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/amc"
   "154.pages.dev/rosso"
   "net/http"
   "os"
   "slices"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      return err
   }
   auth, err := amc.Raw_Auth.Unmarshal(raw)
   if err != nil {
      return err
   }
   raw, err = auth.Refresh()
   if err != nil {
      return err
   }
   os.WriteFile(home + "/amc/auth.json", raw, 0666)
   if !f.s.Info {
      content, err := auth.Content(f.address)
      if err != nil {
         return err
      }
      video, err := content.Video()
      if err != nil {
         return err
      }
      f.s.Name = stream.Name(video)
   }
   play, err := auth.Playback(f.address)
   if err != nil {
      return err
   }
   f.s.Poster = play
   reps, err := func() ([]*dash.Representation, error) {
      s, err := play.HTTP_DASH()
      if err != nil {
         return nil, err
      }
      r, err := http.Get(s.Src)
      if err != nil {
         return nil, err
      }
      defer r.Body.Close()
      
      var media dash.Media
      media.Decode(r.Body)
      return media.Representation("")
   }()
   if err != nil {
      return err
   }
   return f.s.DASH_Sofia(reps, index)
}

func (f flags) login() error {
   raw, err := amc.Unauth()
   if err != nil {
      return err
   }
   auth, err := raw.Unmarshal()
   if err != nil {
      return err
   }
   raw, err = auth.Login(f.email, f.password)
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/amc/auth.json", raw, 0666)
}
