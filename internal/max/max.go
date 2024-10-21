package main

import (
   "41.neocities.org/media/max"
   "41.neocities.org/media/internal"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f *flags) download() error {
   var (
      token max.DefaultToken
      err error
   )
   token.Session.Raw, err = os.ReadFile(f.home + "/session.txt")
   if err != nil {
      return err
   }
   token.Token.Raw, err = os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return err
   }
   err = token.Unmarshal()
   if err != nil {
      return err
   }
   //////////////////////////////////////////////////////////////////////////////
   play, err := token.Playback(f.address)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.Fallback.Manifest.Url.Url, nil)
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
      if rep.GetAdaptationSet().GetPeriod().Id == "0" {
         if rep.GetAdaptationSet().MaxHeight <= f.max_height {
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
   }
   return nil
}

func (f *flags) authenticate() error {
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
   err = token.Unmarshal()
   if err != nil {
      return err
   }
   err = token.Login(key, login)
   if err != nil {
      return err
   }
   os.Mkdir(f.home, os.ModePerm)
   os.WriteFile(f.home + "/session.txt", token.Session.Raw, os.ModePerm)
   return os.WriteFile(f.home + "/token.txt", token.Token.Raw, os.ModePerm)
}
