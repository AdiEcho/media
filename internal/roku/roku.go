package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/roku"
   "fmt"
   "io"
   "net/http"
   "os"
   "sort"
)

func (f *flags) download() error {
   var token *roku.AccountToken
   if f.token_read {
      token = &roku.AccountToken{}
      var err error
      token.Raw, err = os.ReadFile(f.home + "/roku.txt")
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
   resp, err := http.Get(play.Url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
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

func (f *flags) write_token() error {
   var err error
   // AccountAuth
   var auth roku.AccountAuth
   auth.Raw, err = os.ReadFile("auth.txt")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   // AccountCode
   var code roku.AccountCode
   code.Raw, err = os.ReadFile("code.txt")
   if err != nil {
      return err
   }
   code.Unmarshal()
   // AccountToken
   token, err := auth.Token(&code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/roku.txt", token.Raw, os.ModePerm)
}

func write_code() error {
   // AccountAuth
   var auth roku.AccountAuth
   err := auth.New(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("auth.txt", auth.Raw, os.ModePerm)
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
   err = os.WriteFile("code.txt", code.Raw, os.ModePerm)
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
