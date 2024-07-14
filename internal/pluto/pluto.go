package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/pluto"
   "errors"
   "fmt"
   "net/http"
)

func (f flags) download() error {
   video, err := f.address.Video(f.forward)
   if err != nil {
      return err
   }
   clip, err := video.Clip()
   if err != nil {
      return err
   }
   file, ok := clip.Dash()
   if !ok {
      return errors.New("EpisodeClip.Dash")
   }
   req, err := http.NewRequest("", f.base, nil)
   if err != nil {
      return err
   }
   req.URL.Path = file.Path
   media, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         f.s.Name = pluto.Namer{video}
         f.s.Poster = pluto.Poster{}
         return f.s.Download(medium)
      }
   }
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
