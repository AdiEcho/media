package main

import (
   "154.pages.dev/media/hulu"
   "fmt"
   "os"
)

func (f flags) download() error {
   var (
      auth hulu.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/hulu.json")
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
   // 1 MPD one
   media, err := f.h.DashMedia(play.Stream_URL)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         detail, err := auth.Details(deep)
         if err != nil {
            return err
         }
         f.h.Name = <-detail
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
   return os.WriteFile(name, auth.Data, 0666)
}
