package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/pluto"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

func (f flags) download() error {
   video, err := f.web.Video()
   if err != nil {
      return err
   }
   clip, err := video.Clip()
   if err != nil {
      return err
   }
   source, ok := clip.DASH()
   if !ok {
      return errors.New("pluto.EpisodeClip.DASH")
   }
   file, err := source.Parse(f.base)
   if err != nil {
      return err
   }
   // 1 MPD one
   media, err := f.h.DashMedia(file.String())
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         // FIXME
         f.h.Name = meta
         f.h.Poster = pluto.Core()
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
