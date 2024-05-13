package main

import (
   "154.pages.dev/media/cine/member"
   "errors"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   article, err := f.slug.Article()
   if err != nil {
      return err
   }
   var auth member.Authenticate
   auth.Data, err = os.ReadFile(f.home + "cine-member.json")
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   asset, ok := article.Film()
   if !ok {
      return errors.New("member.DataArticle.Film")
   }
   play, err := auth.Play(asset)
   if err != nil {
      return err
   }
   dash, ok := play.DASH()
   if !ok {
      return errors.New("member.AssetPlay.DASH")
   }
   req, err := http.NewRequest("", dash, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Name = member.Namer{article}
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

func (f flags) authenticate() error {
   var auth member.Authenticate
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/cine-member.json", auth.Data, 0666)
}
