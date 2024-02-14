package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/hulu"
   "154.pages.dev/rosso"
   "encoding/xml"
   "net/http"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth hulu.Authenticate
   auth.Raw, err = os.ReadFile(home + "/hulu/token.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.Deep_Link(f.id)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   if !f.s.Info {
      detail, err := auth.Details(deep)
      if err != nil {
         return err
      }
      f.s.Name = rosso.Name(detail)
      f.s.Poster = play
   }
   var media dash.MPD
   err = func() error {
      res, err := http.Get(play.Stream_URL)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      f.s.Base = res.Request.URL
      return xml.NewDecoder(res.Body).Decode(&media)
   }()
   if err != nil {
      return err
   }
   return f.s.DASH_Sofia(media, f.representation)
}

func (f flags) authenticate() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/hulu/token.json"
   auth, err := hulu.Living_Room(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}
