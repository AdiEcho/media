package main

import (
   "154.pages.dev/media/rtbf"
   "154.pages.dev/media/internal"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) authenticate() error {
   var login accounts_login
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   text, err := login.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.json", text, 0666)
}

func (f flags) download() error {
   var (
      auth rtbf.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/rtbf.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.entity)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.StreamUrl, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name, err = auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Poster = play
         return f.s.Download(medium)
      }
   }
   // 2 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}
