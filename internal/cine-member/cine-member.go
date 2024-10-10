package main

import (
   "41.neocities.org/media/cine/member"
   "41.neocities.org/media/internal"
   "errors"
   "fmt"
   "net/http"
   "os"
   "path"
)

func (f *flags) write_play() error {
   os.Mkdir(f.base(), os.ModePerm)
   // 1. write OperationArticle
   article, err := f.address.Article()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.base() + "/article.txt", article.Raw, os.ModePerm)
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
      return article.FilmError()
   }
   var user member.OperationUser
   user.Raw, err = os.ReadFile(f.home + "/cine-member.txt")
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
   return os.WriteFile(f.base() + "/play.txt", play.Raw, os.ModePerm)
}

func (f *flags) write_user() error {
   var user member.OperationUser
   err := user.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/cine-member.txt", user.Raw, os.ModePerm)
}

func (f *flags) download() error {
   var (
      play member.OperationPlay
      err error
   )
   play.Raw, err = os.ReadFile(f.base() + "/play.txt")
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
         article.Raw, err = os.ReadFile(f.base() + "/article.txt")
         if err != nil {
            return err
         }
         err = article.Unmarshal()
         if err != nil {
            return err
         }
         f.s.Name = &article
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) base() string {
   return path.Base(f.address.Path)
}
