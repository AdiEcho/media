package main

import (
   "154.pages.dev/media/ctv"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) base() string {
   return path.Base(string(f.path)) + ".json"
}

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
   text, err := manifest.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.base(), text, 0666)
}

func (f flags) download() error {
   text, err := os.ReadFile(f.base())
   if err != nil {
      return err
   }
   var manifest ctv.MediaManifest
   err = manifest.Unmarshal(text)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", manifest.URL, nil)
   if err != nil {
      return err
   }
   represents, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      if represent.ID == f.representation {
         f.s.Name = manifest
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
