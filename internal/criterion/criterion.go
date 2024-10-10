package main

import (
   "41.neocities.org/media/criterion"
   "41.neocities.org/media/internal"
   "fmt"
   "net/http"
   "os"
   "path"
   "sort"
)

func (f *flags) download() error {
   var (
      token criterion.AuthToken
      err error
   )
   token.Raw, err = os.ReadFile(f.home + "/criterion.txt")
   if err != nil {
      return err
   }
   err = token.Unmarshal()
   if err != nil {
      return err
   }
   item, err := token.Video(path.Base(f.address))
   if err != nil {
      return err
   }
   files, err := token.Files(item)
   if err != nil {
      return err
   }
   file, ok := files.Dash()
   if !ok {
      return files.DashError()
   }
   req, err := http.NewRequest("", file.Links.Source.Href, nil)
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
         if _, ok := rep.Ext(); ok {
            fmt.Print(&rep, "\n\n")
         }
      case rep.Id:
         f.s.Name = item
         f.s.Poster = file
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   var token criterion.AuthToken
   err := token.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.txt", token.Raw, os.ModePerm)
}
