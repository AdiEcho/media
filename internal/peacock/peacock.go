package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/peacock"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

func (f flags) authenticate() error {
   name += "/peacock.json"
   auth, err := peacock.LivingRoom(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth peacock.Authenticate
   auth.Raw, err = os.ReadFile(home + "/peacock.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.peacock_id)
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
