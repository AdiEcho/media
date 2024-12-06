package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/criterion"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
   "sort"
)

func (f *flags) authenticate() error {
   data, err := criterion.AuthToken{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/criterion.txt")
   if err != nil {
      return err
   }
   var token criterion.AuthToken
   err = token.Unmarshal(data)
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
      return errors.New("VideoFiles.Dash")
   }
   resp, err := http.Get(file.Links.Source.Href)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
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
         f.s.Client = file
         return f.s.Download(rep)
      }
   }
   return nil
}
