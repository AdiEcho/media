package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/draken"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
   "sort"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/draken.txt")
   if err != nil {
      return err
   }
   var login draken.AuthLogin
   err = login.Unmarshal(data)
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
   resp, err := http.Get(play.Playlist)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         if _, ok := rep.Ext(); ok {
            fmt.Print(&rep, "\n\n")
         }
      case rep.Id:
         f.s.Namer = &draken.Namer{movie}
         f.s.Wrapper = draken.Wrapper(
            func() (*draken.AuthLogin, *draken.Playback) {
               return &login, play
            },
         )
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := draken.AuthLogin{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/draken.txt", data, os.ModePerm)
}
