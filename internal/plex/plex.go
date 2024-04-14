package main

import (
   "154.pages.dev/media/plex"
   "errors"
   "fmt"
)

func (f flags) download() error {
   var anon plex.Anonymous
   err := anon.New()
   if err != nil {
      return err
   }
   match, err := anon.Discover(f.path)
   if err != nil {
      return err
   }
   video, err := anon.Video(match)
   if err != nil {
      return err
   }
   part, ok := video.DASH(anon)
   if !ok {
      return errors.New("plex.OnDemand.DASH")
   }
   // 1 MPD one
   media, err := f.h.DashMedia(part.Key)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.dash {
         f.h.Poster = part
         f.h.Name = plex.Name{match}
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
