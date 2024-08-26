package main

import (
   "154.pages.dev/media/ctv"
   "154.pages.dev/media/internal"
   "fmt"
   "net/http"
   "os"
   "path"
   "sort"
)

func (f *flags) download() error {
   manifest, err := os.ReadFile(f.base() + "/manifest.txt")
   if err != nil {
      return err
   }
   var media ctv.MediaContent
   media.Raw, err = os.ReadFile(f.base() + "/media.txt")
   if err != nil {
      return err
   }
   err = media.Unmarshal()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", string(manifest), nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = ctv.Namer{&media}
         f.s.Poster = ctv.Poster{}
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) get_manifest() error {
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
   os.Mkdir(f.base(), os.ModePerm)
   err = os.WriteFile(f.base() + "/media.txt", media.Raw, os.ModePerm)
   if err != nil {
      return err
   }
   manifest, err := axis.Manifest(media)
   if err != nil {
      return err
   }
   return os.WriteFile(f.base() + "/manifest.txt", []byte(manifest), os.ModePerm)
}

func (f *flags) base() string {
   return path.Base(string(f.path))
}
