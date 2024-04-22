package main

import (
   "154.pages.dev/media/roku"
   "errors"
   "fmt"
   "net/http"
)

func (f flags) download() error {
   var home roku.HomeScreen
   err := home.New(f.roku)
   if err != nil {
      return err
   }
   video, ok := home.DASH()
   if !ok {
      return errors.New("roku.HomeScreen.DASH")
   }
   req, err := http.NewRequest("", video.URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         var site roku.CrossSite
         err := site.New()
         if err != nil {
            return err
         }
         f.s.Poster, err = site.Playback(f.roku)
         if err != nil {
            return err
         }
         f.s.Name = roku.Namer{home}
         return f.s.Download(medium)
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
