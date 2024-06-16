package main

import (
   "154.pages.dev/media/amc"
   "154.pages.dev/media/internal"
   "errors"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   var (
      auth amc.Authorization
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/amc.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   auth.Refresh()
   os.WriteFile(f.home + "/amc.json", auth.Data, 0666)
   play, err := auth.Playback(f.address.NID)
   if err != nil {
      return err
   }
   source, ok := play.HttpsDash()
   if !ok {
      return errors.New("amc.Playback.HttpsDash")
   }
   req, err := http.NewRequest("", source.Src, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         content, err := auth.Content(f.address.Path)
         if err != nil {
            return err
         }
         video, err := content.Video()
         if err != nil {
            return err
         }
         f.s.Name = video
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

func (f flags) login() error {
   var auth amc.Authorization
   auth.Unauth()
   auth.Unmarshal()
   auth.Login(f.email, f.password)
   return os.WriteFile(f.home + "/amc.json", auth.Data, 0666)
}
