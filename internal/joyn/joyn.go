package main

import (
   "154.pages.dev/media/joyn"
   "errors"
   "fmt"
   "net/http"
   "net/url"
)

func (f flags) download() error {
   address, err := url.Parse(f.address)
   if err != nil {
      return err
   }
   detail, err := joyn.NewDetail(address.Path)
   if err != nil {
      return err
   }
   var anonymous joyn.Anonymous
   err = anonymous.New()
   if err != nil {
      return err
   }
   content_id, ok := detail.ContentId()
   if !ok {
      return errors.New("joyn.DetailPage.ContentId")
   }
   title, err := anonymous.Entitlement(content_id)
   if err != nil {
      return err
   }
   play, err := title.Playlist(content_id)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.ManifestUrl, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name = joyn.Namer{detail}
         f.s.Poster = play
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
