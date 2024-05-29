package main

import (
   "154.pages.dev/media/roku"
   "fmt"
   "net/http"
)

func (f flags) get_code() error {
   var token roku.AccountToken
   token.New(nil)
   // FIXME
   code, err := token.code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
   text, err := code.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", text, 0666)
}

func (f flags) download() error {
   var token roku.AccountToken
   err := token.New(nil)
   if err != nil {
      return err
   }
   play, err := token.Playback(f.roku)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", play.URL, nil)
   if err != nil {
      return err
   }
   media, err := f.s.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         var home roku.HomeScreen
         err := home.New(f.roku)
         if err != nil {
            return err
         }
         f.s.Name = roku.Namer{home}
         f.s.Poster = play
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
