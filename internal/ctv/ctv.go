package main

import (
   "154.pages.dev/media/ctv"
   "fmt"
   "net/http"
)

func (f flags) get_manifest() error {
   resolve, err := f.path.Resolve()
   if err != nil {
      return err
   }
   axis, err := resolve.Axis()
   if err != nil {
      return err
   }
   media, err := axis.Media()
   if err != nil {
      return err
   }
   manifest, err := axis.Manifest(media)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", manifest.URL, nil)
   if err != nil {
      return err
   }
   represents, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      if represent.ID == f.representation {
         f.s.Name = ctv.Namer{media}
         f.s.Poster = ctv.Poster{}
         return f.s.Download(represent)
      }
   }
   // 2 MPD all
   for i, represent := range represents {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(represent)
   }
   return nil
}

func (f flags) download() error {
   resolve, err := f.path.Resolve()
   if err != nil {
      return err
   }
   axis, err := resolve.Axis()
   if err != nil {
      return err
   }
   media, err := axis.Media()
   if err != nil {
      return err
   }
   manifest, err := axis.Manifest(media)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", manifest.URL, nil)
   if err != nil {
      return err
   }
   represents, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      if represent.ID == f.representation {
         f.s.Name = ctv.Namer{media}
         f.s.Poster = ctv.Poster{}
         return f.s.Download(represent)
      }
   }
   // 2 MPD all
   for i, represent := range represents {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(represent)
   }
   return nil
}
