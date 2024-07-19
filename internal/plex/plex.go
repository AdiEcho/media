package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "fmt"
   "net/http"
)

func (f flags) download() error {
   var anon plex.Anonymous
   err := anon.New()
   if err != nil {
      return err
   }
   match, err := anon.Discover(f.address)
   if err != nil {
      return err
   }
   video, err := anon.Video(match, f.forward)
   if err != nil {
      return err
   }
   part, ok := video.Dash(anon)
   if !ok {
      return plex.MediaPart{}
   }
   req, err := http.NewRequest("", part.Key, nil)
   if err != nil {
      return err
   }
   if f.forward != "" {
      req.Header.Set("x-forwarded-for", f.forward)
   }
   media, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         f.s.Poster = part
         f.s.Name = plex.Namer{match}
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
