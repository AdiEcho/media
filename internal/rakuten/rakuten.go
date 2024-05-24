package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/rakuten"
   "154.pages.dev/media/internal"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) write_stream() error {
   fhd, err := web.FHD().Info()
   if err != nil {
      return err
   }
   hd, err := web.HD().Info()
   if err != nil {
      return err
   }
   fhd.LicenseUrl = hd.LicenseUrl
   text, err := fhd.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.name(), text, 0666)
}

func (f flags) name() string {
   return path.Base(f.web.String()) + ".json"
}

func (f flags) download() error {
   text, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   var info rakuten.StreamInfo
   err = info.Unmarshal(text)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", info.URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Poster = info
         
         // FIXME
         detail, err := auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Name = <-detail
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
