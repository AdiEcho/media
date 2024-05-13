package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/cine/member"
   "154.pages.dev/media/internal"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
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
   
   play, err := auth.play(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.dash())
   req, err := http.NewRequest("", play.Stream_URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         detail, err := auth.Details(deep)
         if err != nil {
            return err
         }
         f.s.Name = <-detail
         f.s.Poster = play
         return f.s.Download(medium)
      }
   }
   // 2 MPD all
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
