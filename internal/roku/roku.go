package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/roku"
   "fmt"
   "net/http"
   "os"
)

func (f flags) write_code() error {
   os.Mkdir(f.roku, 0666)
   // AccountToken
   var account roku.AccountToken
   err := account.New(nil)
   if err != nil {
      return err
   }
   err = account.Unmarshal()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.roku + "/account.json", account.Data, 0666)
   if err != nil {
      return err
   }
   // ActivationCode
   code, err := account.Code()
   if err != nil {
      return err
   }
   err = code.Unmarshal()
   if err != nil {
      return err
   }
   fmt.Println(code)
   return os.WriteFile(f.roku + "/code.json", code.Data, 0666)
}

func (f flags) write_token() error {
   var err error
   // AccountToken
   var account roku.AccountToken
   account.Data, err = os.ReadFile(f.roku + "/account.json")
   if err != nil {
      return err
   }
   account.Unmarshal()
   // ActivationCode
   var code roku.ActivationCode
   code.Data, err = os.ReadFile(f.roku + "/code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   // ActivationToken
   activation_token, err := account.ActivationToken(code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/roku.json", activation_token.Data, 0666)
}

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
   req, err := http.NewRequest("", play.Url, nil)
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
         var home roku.HomeScreen
         err := home.New(f.roku)
         if err != nil {
            return err
         }
         f.s.Name = roku.Namer{home}
         f.s.Poster = play
         return f.s.Download(rep)
      }
   }
   return nil
}
