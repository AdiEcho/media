package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/plex"
   "errors"
   "fmt"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   var user plex.Anonymous
   err := user.New()
   if err != nil {
      return err
   }
   match, err := user.Match(f.address)
   if err != nil {
      return err
   }
   video, err := user.Video(match, f.set_forward)
   if err != nil {
      return err
   }
   part, ok := video.Dash()
   if !ok {
      return errors.New("OnDemand.Dash")
   }
   req, err := http.NewRequest("", part.Key.Url.String(), nil)
   if err != nil {
      return err
   }
   if f.set_forward != "" {
      req.Header.Set("x-forwarded-for", f.set_forward)
   }
   reps, err := internal.Mpd(req)
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

func get_forward() {
   for _, forward := range internal.Forward {
      fmt.Println(forward.Country, forward.IP)
   }
}
