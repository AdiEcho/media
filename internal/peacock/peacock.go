package main

import (
   "154.pages.dev/media/peacock"
   "errors"
   "fmt"
   "os"
)

func (f flags) authenticate() error {
   var session peacock.IdSession
   session.New(f.id_session)
   text, err := session.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/peacock.json", text, 0666)
}

func (f flags) download() error {
   text, err := os.ReadFile(f.home + "/peacock.json")
   if err != nil {
      return err
   }
   var session peacock.IdSession
   session.Unmarshal(text)
   auth, err := session.Auth()
   if err != nil {
      return err
   }
   video, err := auth.Video(f.peacock_id)
   if err != nil {
      return err
   }
   akamai, ok := video.Akamai()
   if !ok {
      return errors.New("peacock.VideoPlayout.Akamai")
   }
   // 1 MPD one
   media, err := f.h.DashMedia(akamai)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         var node peacock.QueryNode
         if err := node.New(f.peacock_id); err != nil {
            return err
         }
         f.h.Name = node
         f.h.Poster = video
         return f.h.DASH(medium)
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
