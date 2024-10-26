package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/pluto"
   "errors"
   "fmt"
   "net/http"
   "sort"
)

func (f *flags) download() error {
   video, err := f.address.Video(f.set_forward)
   if err != nil {
      return err
   }
   clip, err := video.Clip()
   if err != nil {
      return err
   }
   var (
      req http.Request
      ok bool
   )
   req.URL, ok = clip.Dash()
   if !ok {
      return errors.New("EpisodeClip.Dash")
   }
   req.URL.Scheme = pluto.Base[0].Scheme
   req.URL.Host = pluto.Base[0].Host
   reps, err := internal.Mpd(&req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Name = pluto.Namer{video}
         f.s.Poster = pluto.Poster{}
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
