package main

import (
   "154.pages.dev/media/cine/member"
   "154.pages.dev/media/internal"
   "errors"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f flags) write_play() error {
   os.Mkdir(f.base(), 0666)
   // 1. write OperationArticle
   article, err := f.slug.Article()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.base() + "/article.json", article.Data, 0666)
   if err != nil {
      return err
   }
   err = article.Unmarshal()
   if err != nil {
      return err
   }
   // 2. write OperationPlay
   asset, ok := article.Film()
   if !ok {
      return member.ArticleAsset{}
   }
   var user member.OperationUser
   user.Data, err = os.ReadFile(f.home + "/cine-member.json")
   if err != nil {
      return err
   }
   err = user.Unmarshal()
   if err != nil {
      return err
   }
   play, err := user.Play(asset)
   if err != nil {
      return err
   }
   return os.WriteFile(f.base() + "/play.json", play.Data, 0666)
}

func (f flags) write_user() error {
   var user member.OperationUser
   err := user.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/cine-member.json", user.Data, 0666)
}

func (f flags) base() string {
   return path.Base(string(f.slug))
}

func (f flags) download() error {
   var (
      play member.OperationPlay
      err error
   )
   play.Data, err = os.ReadFile(f.base() + "/play.json")
   if err != nil {
      return err
   }
   err = play.Unmarshal()
   if err != nil {
      return err
   }
   dash, ok := play.Dash()
   if !ok {
      return errors.New("OperationPlay.Dash")
   }
   req, err := http.NewRequest("", dash, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         var article member.OperationArticle
         article.Data, err = os.ReadFile(f.base() + "/article.json")
         if err != nil {
            return err
         }
         err = article.Unmarshal()
         if err != nil {
            return err
         }
         f.s.Name = article
         return f.s.Download(rep)
      }
   }
   return nil
}
