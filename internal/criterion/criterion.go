package main

import (
   "154.pages.dev/media/criterion"
   "154.pages.dev/media/internal"
   "errors"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) download() error {
   var (
      token criterion.AuthToken
      err error
   )
   token.Data, err = os.ReadFile(f.home + "/criterion.json")
   if err != nil {
      return err
   }
   token.Unmarshal()
   item, err := token.Video(path.Base(f.address))
   if err != nil {
      return err
   }
   files, err := token.Files(item)
   if err != nil {
      return err
   }
   file, ok := files.DASH()
   if !ok {
      return errors.New("criterion.VideoFiles.DASH")
   }
   req, err := http.NewRequest("", file.Links.Source.Href, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name = item
         f.s.Poster = file
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

func (f flags) authenticate() error {
   var token criterion.AuthToken
   err := token.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.json", token.Data, 0666)
}
