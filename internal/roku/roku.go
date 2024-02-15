package main

import "154.pages.dev/media/roku"

func (f flags) DASH(home roku.HomeScreen) error {
   if f.dash_id != "" {
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
   }
   video, ok := home.DASH()
   if ok {
      media, err := f.h.DashMedia(video.URL)
      if err != nil {
         return err
      }
      return f.h.DASH(media, f.dash_id)
   }
   return nil
}
