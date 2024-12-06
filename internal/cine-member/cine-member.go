package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/cine/member"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
)

func (f *flags) write_user() error {
   data, err := member.OperationUser{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/cine-member.txt", data, os.ModePerm)
}

func (f *flags) write_play() error {
   os.Mkdir(f.base(), os.ModePerm)
   // 1. write OperationArticle
   var article member.OperationArticle
   data, err := article.Marshal(&f.address)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.base() + "/article.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = article.Unmarshal(data)
   if err != nil {
      return err
   }
   // 2. write OperationPlay
   data, err = os.ReadFile(f.home + "/cine-member.txt")
   if err != nil {
      return err
   }
   var user member.OperationUser
   err = user.Unmarshal(data)
   if err != nil {
      return err
   }
   asset, ok := article.Film()
   if !ok {
      return errors.New("OperationArticle.Film")
   }
   data, err = member.OperationPlay{}.Marshal(&user, asset)
   if err != nil {
      return err
   }
   return os.WriteFile(f.base() + "/play.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.base() + "/play.txt")
   if err != nil {
      return err
   }
   var play member.OperationPlay
   err = play.Unmarshal(data)
   if err != nil {
      return err
   }
   title, ok := play.Dash()
   if !ok {
      return errors.New("OperationPlay.Dash")
   }
   resp, err := http.Get(title.Manifest)
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
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         data, err = os.ReadFile(f.base() + "/article.txt")
         if err != nil {
            return err
         }
         var article member.OperationArticle
         err = article.Unmarshal(data)
         if err != nil {
            return err
         }
         f.s.Name = &article
         f.s.Client = title
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) base() string {
   return path.Base(f.address.Path)
}
