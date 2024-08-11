package main

import (
   "154.pages.dev/media/amc"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   var (
      auth amc.Authorization
      err error
   )
   auth.Raw, err = os.ReadFile(f.home + "/amc.txt")
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   err = auth.Refresh()
   if err != nil {
      return err
   }
   os.WriteFile(f.home + "/amc.txt", auth.Raw, 0666)
   
   err = auth.Unmarshal(raw)
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.address.Nid)
   if err != nil {
      return err
   }
   source, ok := play.HttpsDash()
   if !ok {
      return amc.DataSource{}
   }
   req, err := http.NewRequest("", source.Src, nil)
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
         f.s.Poster = play
         content, err := auth.Content(f.address.Path)
         if err != nil {
            return err
         }
         f.s.Name, err = content.Video()
         if err != nil {
            return err
         }
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) login() error {
   var auth amc.Authorization
   err := auth.Unauth()
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   err = auth.Login(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/amc.txt", auth.Raw, 0666)
}
