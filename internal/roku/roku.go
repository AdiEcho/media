package main

import (
   "154.pages.dev/media/roku"
   "errors"
   "fmt"
)

func (f flags) download() error {
   var home roku.HomeScreen
   err := home.New(f.roku_id)
   if err != nil {
      return err
   }
   video, ok := home.DASH()
   if !ok {
      return errors.New("roku.HomeScreen.DASH")
   }
   // 1 MPD one
   media, err := f.h.DashMedia(video.URL)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         var site roku.CrossSite
         err := site.New()
         if err != nil {
            return err
         }
         f.h.Poster, err = site.Playback(f.roku_id)
         if err != nil {
            return err
         }
         f.h.Name = home
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
