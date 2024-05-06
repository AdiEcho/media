package main

import (
   "154.pages.dev/media/draken"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) download() error {
   var (
      auth draken.AuthLogin
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/draken.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   movie, err := draken.NewMovie(path.Base(f.address))
   if err != nil {
      return err
   }
   title, err := auth.Entitlement(movie)
   if err != nil {
      return err
   }
   play, err := auth.Playback(movie, title)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.Playlist, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name = draken.Namer{movie}
         f.s.Poster = draken.Poster{auth, play}
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
   var auth draken.AuthLogin
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/draken.json", auth.Data, 0666)
}
