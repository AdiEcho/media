package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/plex"
   "errors"
   "fmt"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   var anon plex.Anonymous
   err := anon.New()
   if err != nil {
      return err
   }
   match, err := anon.Discover(f.url)
   if err != nil {
      return err
   }
   video, err := anon.Video(match, f.forward)
   if err != nil {
      return err
   }
   part, ok := video.Dash(anon)
   if !ok {
      return errors.New("OnDemand.Dash")
   }
   req, err := http.NewRequest("", part.Key, nil)
   if err != nil {
      return err
   }
   if f.forward != "" {
      req.Header.Set("x-forwarded-for", f.forward)
   }
   reps, err := internal.Dash(req)
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
            fmt.Print(rep, "\n\n")
         }
      case rep.Id:
         f.s.Name = &plex.Namer{match}
         f.s.Poster = part
         return f.s.Download(rep)
      }
   }
   return nil
}
