package main

import (
   "154.pages.dev/media/roku"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   var activate *roku.ActivationToken
   if f.token_read {
      activate = new(roku.ActivationToken)
      var err error
      activate.Data, err = os.ReadFile(f.home + "/roku.json")
      if err != nil {
         return err
      }
      activate.Unmarshal()
   }
   var account roku.AccountToken
   account.New(activate)
   play, err := account.Playback(f.roku)
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

func (f flags) write_token() error {
   text, err := os.ReadFile("code.json")
   if err != nil {
      return err
   }
   var code roku.ActivationCode
   code.Unmarshal(text)
   token, err := code.Token()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/roku.json", token.Data, 0666)
}

func write_code() error {
   var token roku.AccountToken
   token.New(nil)
   code, err := token.Code()
   if err != nil {
      return err
   }
   fmt.Println(code)
   text, err := code.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile("code.json", text, 0666)
}
