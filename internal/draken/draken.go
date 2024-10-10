package main

import (
   "41.neocities.org/media/draken"
   "41.neocities.org/media/internal"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f *flags) download() error {
   var (
      login draken.AuthLogin
      err error
   )
   login.Raw, err = os.ReadFile(f.home + "/draken.txt")
   if err != nil {
      return err
   }
   err = login.Unmarshal()
   if err != nil {
      return err
   }
   var movie draken.FullMovie
   err = movie.New(path.Base(f.address))
   if err != nil {
      return err
   }
   title, err := login.Entitlement(&movie)
   if err != nil {
      return err
   }
   play, err := login.Playback(&movie, title)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.Playlist, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = &draken.Namer{&movie}
         f.s.Poster = &draken.Poster{&login, play}
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   var login draken.AuthLogin
   err := login.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/draken.txt", login.Raw, os.ModePerm)
}
