package main

import (
   "154.pages.dev/media/hulu"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f flags) download() error {
   var (
      auth hulu.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/hulu.json")
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
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      if rep.GetAdaptationSet().GetPeriod().Id == "content-0" {
         switch f.representation {
         case "":
            fmt.Print(rep, "\n\n")
         case rep.Id:
            f.s.Name, err = auth.Details(deep)
            if err != nil {
               return err
            }
            f.s.Poster = play
            return f.s.Download(rep)
         }
      }
   }
   return nil
}

func (f flags) authenticate() error {
   var auth hulu.Authenticate
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/hulu.json", auth.Data, 0666)
}
