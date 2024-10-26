package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/itv"
   "errors"
   "fmt"
   "net/http"
   "path"
)

func (f *flags) download() error {
   var id itv.LegacyId
   err := id.Set(path.Base(f.address))
   if err != nil {
      return err
   }
   discovery, err := id.Discovery()
   if err != nil {
      return err
   }
   play, err := discovery.Playlist()
   if err != nil {
      return err
   }
   address, ok := play.Resolution720()
   if !ok {
      return errors.New("resolution 720")
   }
   req, err := http.NewRequest("", address, nil)
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
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Name = itv.Namer{discovery}
         f.s.Poster = itv.Poster{}
         return f.s.Download(rep)
      }
   }
   return nil
}
