package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/max"
   "fmt"
   "net/http"
   "os"
   "sort"
   "testing"
   "time"
)

func (f *flags) do_initiate() error {
   var token max.BoltToken
   err := token.New()
   if err != nil {
      return err
   }
   os.WriteFile("token.txt", []byte(token.St), os.ModePerm)
   initiate, err := token.Initiate()
   if err != nil {
      return err
   }
   fmt.Printf("%+v\n", initiate)
   return nil
}

func (f *flags) do_login() error {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      return err
   }
   var token max.BoltToken
   token.St = string(data)
   login, err := token.Login()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.home + "/state.txt", []byte(login.State), os.ModePerm)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/login.txt", login.Raw, os.ModePerm)
}

func (f *flags) download() error {
   var login max.LinkLogin
   // head
   state, err := os.ReadFile(f.home + "/state.txt")
   if err != nil {
      return err
   }
   login.State = string(state)
   // body
   login.Raw, err = os.ReadFile(f.home + "/login.txt")
   if err != nil {
      return err
   }
   err = login.Unmarshal()
   if err != nil {
      return err
   }
   play, err := login.Playback(f.address)
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
   for _, rep := range reps {
      if rep.GetAdaptationSet().GetPeriod().Id == "0" {
         switch f.representation {
         case "":
            fmt.Print(rep, "\n\n")
         case rep.Id:
            f.s.Name, err = login.Routes(f.address)
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
