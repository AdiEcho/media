package main

import (
   "154.pages.dev/media/amc"
   "os"
)

func (f flags) login() error {
   var auth amc.Authorization
   auth.Unauth()
   auth.Unmarshal()
   auth.Login(f.email, f.password)
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/amc.json", auth.Raw, 0666)
}

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth amc.Authorization
   auth.Raw, err = os.ReadFile(home + "/amc.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   auth.Refresh()
   os.WriteFile(home + "/amc.json", auth.Raw, 0666)
   if f.dash_id != "" {
      content, err := auth.Content(f.web.Path)
      if err != nil {
         return err
      }
      video, err := content.Video()
      if err != nil {
         return err
      }
      f.h.Name = video
   }
   play, err := auth.Playback(f.web.NID)
   if err != nil {
      return err
   }
   f.h.Poster = play
   source, ok := play.HttpsDash()
   if ok {
      media, err := f.h.DashMedia(source.Src)
      if err != nil {
         return err
      }
      return f.h.DASH(media, f.dash_id)
   }
   return nil
}
