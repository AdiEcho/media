package main

import (
   "154.pages.dev/media/tubi"
   "errors"
   "fmt"
   "net/http"
)

func (f flags) write_content() error {
   content := new(tubi.Content)
   err := content.New(f.tubi)
   if err != nil {
      return err
   }
   
   
   
   if content.EpisodeType() {
      err := content.New(content.V.SeriesId)
      if err != nil {
         return err
      }
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("tubi.Content.Get")
      }
   }
   
   
   
}

func (f flags) download() error {
   content := new(tubi.Content)
   err := content.New(f.tubi)
   if err != nil {
      return err
   }
   if content.EpisodeType() {
      err := content.New(content.V.SeriesId)
      if err != nil {
         return err
      }
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("tubi.Content.Get")
      }
   }
   video, err := content.Video()
   if err != nil {
      return err
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
         f.s.Name = content
         f.s.Poster = video
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
