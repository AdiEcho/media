package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/joyn"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) download() error {
   var anon anonymous
   err = anon.New()
   if err != nil {
      t.Fatal(err)
   }
   detail, err := new_detail(test.path)
   if err != nil {
      t.Fatal(err)
   }
   content_id, ok := detail.content_id()
   if !ok {
      t.Fatal("detail_page.content_id")
   }
   title, err := anon.entitlement(content_id)
   if err != nil {
      t.Fatal(err)
   }
   play, err := title.playlist(content_id)
   if err != nil {
      t.Fatal(err)
   }
   req, err := http.NewRequest("", demand.PlaybackUrl, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name = meta
         f.s.Poster = joyn.Core()
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
