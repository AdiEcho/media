package main

import (
   "154.pages.dev/media/max"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) download() error {
   var (
      auth max.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/max.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.entity)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.StreamUrl, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name, err = auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Poster = play
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

func (f flags) authenticate() error {
   var auth max.Authenticate
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/max.json", auth.Data, 0666)
}
