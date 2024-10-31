package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/amc"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "sort"
)

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
   return os.WriteFile(f.home + "/amc.txt", auth.Raw, os.ModePerm)
}

func (f *flags) download() error {
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
   os.WriteFile(f.home + "/amc.txt", auth.Raw, os.ModePerm)
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.address.Nid)
   if err != nil {
      return err
   }
   source, ok := play.Dash()
   if !ok {
      return errors.New("Playback.Dash")
   }
   resp, err := http.Get(source.Src)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
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
         f.s.Poster = play
         content, err := auth.Content(f.address.Path)
         if err != nil {
            return err
         }
         f.s.Name, ok = content.Video()
         if !ok {
            return errors.New("ContentCompiler.Video")
         }
         return f.s.Download(rep)
      }
   }
   return nil
}
