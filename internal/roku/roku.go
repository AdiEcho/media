package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/roku"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f flags) download() error {
   var token *roku.AccountToken
   if f.token_read {
      token = &roku.AccountToken{}
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
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
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

func (f flags) write_token() error {
   var err error
   // AccountAuth
   var auth roku.AccountAuth
   auth.Data, err = os.ReadFile("auth.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   // AccountCode
   var code roku.AccountCode
   code.Data, err = os.ReadFile("code.json")
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

func (f flags) write_code() error {
   // AccountAuth
   var auth roku.AccountAuth
   err := auth.New(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("auth.json", auth.Data, 0666)
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
   err = os.WriteFile("code.json", code.Data, 0666)
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
