package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/rakuten"
   "154.pages.dev/media/internal"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) write_stream() error {
   stream, err := web.fhd().stream()
   // OLD
   var auth rakuten.Authenticate
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/rakuten.json", auth.Data, 0666)
}

func (f flags) download() error {
   var (
      auth rakuten.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/rakuten.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.rakuten)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.Stream_URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         detail, err := auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Name = <-detail
         f.s.Poster = play
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
