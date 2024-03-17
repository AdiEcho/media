package main

import (
   "154.pages.dev/media/amc"
   "errors"
   "fmt"
   "os"
)

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
   play, err := auth.Playback(f.web.NID)
   if err != nil {
      return err
   }
   source, ok := play.HttpsDash()
   if !ok {
      return errors.New("amc.Playback.HttpsDash")
   }
   // 1 MPD one
   media, err := f.h.DashMedia(source.Src)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         content, err := auth.Content(f.web.Path)
         if err != nil {
            return err
         }
         video, err := content.Video()
         if err != nil {
            return err
         }
         f.h.Name = video
         f.h.Poster = play
         return f.h.DASH(medium)
      }
   }
   // 2 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}

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
