package main

import (
   "154.pages.dev/media/max"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f flags) download() error {
   text, err := os.ReadFile(f.home + "/max.json")
   if err != nil {
      return err
   }
   var token max.DefaultToken
   err = token.Unmarshal(text)
   if err != nil {
      return err
   }
   play, err := token.Playback(f.address)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.Manifest.Url, nil)
   if err != nil {
      return err
   }
   reps, err := internal.DASH(req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      if rep.GetAdaptationSet().GetPeriod().Id == "2" {
         switch f.representation {
         case "":
            if _, ok := rep.Ext(); ok {
               fmt.Print(rep, "\n\n")
            }
         case rep.Id:
            f.s.Name, err = token.Routes(f.address)
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
   var login max.DefaultLogin
   login.Credentials.Username = f.email
   login.Credentials.Password = f.password
   var key max.PublicKey
   err := key.New()
   if err != nil {
      return err
   }
   var token max.DefaultToken
   err = token.New()
   if err != nil {
      return err
   }
   err = token.Login(key, login)
   if err != nil {
      return err
   }
   text, err := token.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/max.json", text, 0666)
}
