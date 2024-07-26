package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/roku"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   var token *roku.AccountToken
   if f.token_read {
      token = new(roku.AccountToken)
      var err error
      token.Data, err = os.ReadFile(f.home + "/roku.json")
      if err != nil {
         return err
      }
      token.Unmarshal()
   }
   var auth roku.AccountAuth
   err := auth.New(token)
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.roku)
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

func (f flags) write_code() error {
   os.Mkdir(f.roku, 0666)
   // AccountAuth
   var auth roku.AccountAuth
   err := auth.New(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.roku + "/auth.json", auth.Data, 0666)
   if err != nil {
      return err
   }
   err = auth.Unmarshal()
   if err != nil {
      return err
   }
   // AccountCode
   code, err := auth.Code()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.roku + "/code.json", code.Data, 0666)
   if err != nil {
      return err
   }
   err = code.Unmarshal()
   if err != nil {
      return err
   }
   fmt.Println(code)
   return nil
}

func (f flags) write_token() error {
   var err error
   // AccountAuth
   var auth roku.AccountAuth
   auth.Data, err = os.ReadFile(f.roku + "/auth.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   // AccountCode
   var code roku.AccountCode
   code.Data, err = os.ReadFile(f.roku + "/code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   // AccountToken
   token, err := auth.Token(code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/roku.json", token.Data, 0666)
}
