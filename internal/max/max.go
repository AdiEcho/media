package main

import (
   "154.pages.dev/media/max"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
)

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
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Poster = play
         f.s.Name, err = token.Routes(f.address)
         if err != nil {
            return err
         }
         return f.s.Download(medium)
      }
   }
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
