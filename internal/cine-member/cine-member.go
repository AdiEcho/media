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

func (f flags) play_write() error {
   article, err := f.slug.Article()
   if err != nil {
      return err
   }
   asset, ok := article.Film()
   if !ok {
      return errors.New("member.DataArticle.Film")
   }
   var auth member.Authenticate
   auth.Data, err = os.ReadFile(f.home + "/cine-member.json")
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   play, err := auth.Play(asset)
   if err != nil {
      return err
   }
   text, err := member.Encoding{article, play}.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.play_name(), text, 0666)
}

func (f flags) download() error {
   text, err := os.ReadFile(f.play_name())
   if err != nil {
      return err
   }
   var encode member.Encoding
   err = encode.Unmarshal(text)
   if err != nil {
      return err
   }
   dash, ok := encode.Play.Dash()
   if !ok {
      return errors.New("AssetPlay.Dash")
   }
   req, err := http.NewRequest("", dash, nil)
   if err != nil {
      return err
   }
   media, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         f.s.Name = encode
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

func (f flags) play_name() string {
   return path.Base(string(f.slug)) + ".json"
}
