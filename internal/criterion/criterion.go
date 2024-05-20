package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/criterion"
   "154.pages.dev/media/internal"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) authenticate() error {
   var token criterion.AuthToken
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.json", auth.Data, 0666)
}

func (f flags) download() error {
   var (
      token criterion.AuthToken
      err error
   )
   token.Data, err = os.ReadFile("token.json")
   if err != nil {
      return err
   }
   token.Unmarshal()
   item, err := token.Video(path.Base(f.address))
   if err != nil {
      return err
   }
   
   files, err := token.files(item)
   if err != nil {
      t.Fatal(err)
   }
   file, ok := files.dash()
   if !ok {
      t.Fatal("video_files.dash")
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
         detail, err := token.Details(deep)
         if err != nil {
            return err
         }
         f.s.Name = <-detail
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

