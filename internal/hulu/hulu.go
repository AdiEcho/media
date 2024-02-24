package main

import (
   "154.pages.dev/media/hulu"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth hulu.Authenticate
   auth.Raw, err = os.ReadFile(home + "/hulu.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.hulu_id)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      detail, err := auth.Details(deep)
      if err != nil {
         return err
      }
      f.h.Name = detail
      f.h.Poster = play
   }
   media, err := f.h.DashMedia(play.Stream_URL)
   if err != nil {
      return err
   }
   return f.h.DASH(media, f.dash_id)
}

func (f flags) authenticate() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/hulu.json"
   auth, err := hulu.LivingRoom(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}
