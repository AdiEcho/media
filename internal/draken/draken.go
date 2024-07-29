package main

import (
   "154.pages.dev/media/draken"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) authenticate() error {
   var login draken.AuthLogin
   err := login.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/draken.json", login.Marshal(), 0666)
}

func (f flags) download() error {
   raw, err := os.ReadFile(f.home + "/draken.json")
   if err != nil {
      return err
   }
   var login draken.AuthLogin
   err = login.Unmarshal(raw)
   if err != nil {
      return err
   }
   movie, err := draken.NewMovie(path.Base(f.address))
   if err != nil {
      return err
   }
   title, err := login.Entitlement(movie)
   if err != nil {
      return err
   }
   play, err := login.Playback(movie, title)
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
         f.s.Name = draken.Namer{movie}
         f.s.Poster = draken.Poster{login, play}
         return f.s.Download(rep)
      }
   }
   return nil
}
