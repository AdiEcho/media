package main

import (
   "154.pages.dev/media/tubi"
   "errors"
   "fmt"
   "net/http"
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
   video, ok := content.Video()
   if !ok {
      return errors.New("tubi.Content.Video")
   }
   req, err := http.NewRequest("", video.Manifest.URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Poster = video
         f.s.Name = tubi.Namer{content}
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
