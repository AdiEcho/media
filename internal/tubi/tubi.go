package main

import (
   "154.pages.dev/media/tubi"
   "errors"
   "fmt"
)

func (f flags) download() error {
   content := new(tubi.Content)
   err := content.New(f.tubi)
   if err != nil {
      return err
   }
   if content.Episode() {
      err := content.New(content.Series_ID)
      if err != nil {
         return err
      }
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("tubi.Content.Get")
      }
   }
   video := content.Video()
   // 1 MPD one
   media, err := f.h.DashMedia(video.Manifest.URL)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.h.Poster = video
         f.h.Name = tubi.Namer{content}
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
