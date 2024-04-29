package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/ctv"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

func (f flags) download() error {
   resolve, err := new_resolve(path)
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(99 * time.Millisecond)
   axis, err := resolve.axis()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(99 * time.Millisecond)
   media, err := axis.media()
   if err != nil {
      t.Fatal(err)
   }
   manifest, err := axis.manifest(media)
   if err != nil {
      t.Fatal(err)
   }
   // old
   var meta ctv.Metadata
   err := meta.New(f.ctv)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
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
         f.s.Poster = ctv.Core()
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
