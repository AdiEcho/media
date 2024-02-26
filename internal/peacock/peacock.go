package main

import (
   "154.pages.dev/media/peacock"
   "errors"
   "os"
)

func (f flags) download() error {
   text, err := os.ReadFile(f.home + "peacock.json")
   if err != nil {
      return err
   }
   var sign peacock.SignIn
   sign.Unmarshal(text)
   auth, err := sign.Auth()
   if err != nil {
      return err
   }
   video, err := auth.Video(f.peacock_id)
   if err != nil {
      return err
   }
   var node peacock.QueryNode
   if err := node.New(f.peacock_id); err != nil {
      return err
   }
   if f.dash_id != "" {
      f.h.Name = node
      f.h.Poster = video
   }
   akamai, ok := video.Akamai()
   if !ok {
      return errors.New("peacock.VideoPlayout.Akamai")
   }
   media, err := f.h.DashMedia(akamai)
   if err != nil {
      return err
   }
   return f.h.DASH(media, f.dash_id)
}

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
