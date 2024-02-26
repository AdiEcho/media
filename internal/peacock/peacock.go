package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/peacock"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

func (f flags) authenticate() error {
   var sign peacock.SignIn
   err := sign.New(f.email, f.password)
   if err != nil {
      return err
   }
   text, err := sign.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "peacock.json", text, 0666)
}

func (f flags) download() error {
   text, err := os.ReadFile(f.home + "peacock.json")
   if err != nil {
      return err
   }
   var sign peacock.SignIn
   sign.Unmarshal(text)
   
   auth, err := sign.auth()
   if err != nil {
      t.Fatal(err)
   }
   video, err := auth.video(content_id)
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(video)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Println(key, ok)
   // Namer
   var node query_node
   err := node.New(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(encoding.Name(node))
   // OLD
   var auth peacock.Authenticate
   auth.Raw, err = os.ReadFile(home + "/peacock.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.peacock_id)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      detail, err := auth.Details(deep)
      if err != nil {
         return err
      }
      f.h.Name = detail
      f.h.Poster = play
   }
   media, err := f.h.DashMedia(play.Stream_URL)
   if err != nil {
      return err
   }
   return f.h.DASH(media, f.dash_id)
}
