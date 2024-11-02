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
      data, err := os.ReadFile(f.home + "/roku.txt")
      if err != nil {
         return err
      }
      token = &roku.AccountToken{}
      err = token.Unmarshal(data)
      if err != nil {
         return err
      }
   }
   var auth roku.AccountAuth
   err := auth.New(token, nil)
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

func write_code() error {
   // AccountAuth
   data, err := (*roku.AccountAuth).Marshal(nil, nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("auth.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // AccountCode
   var auth roku.AccountAuth
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = (*roku.AccountCode).Marshal(nil, &auth)
   if err != nil {
      return err
   }
   err = os.WriteFile("code.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   var code roku.AccountCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(code)
   return nil
}
func (f *flags) write_token() error {
   // AccountAuth
   data, err := os.ReadFile("auth.txt")
   if err != nil {
      return err
   }
   var auth roku.AccountAuth
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   // AccountCode
   data, err = os.ReadFile("code.txt")
   if err != nil {
      return err
   }
   var code roku.AccountCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   // AccountToken
   data, err = (*roku.AccountToken).Marshal(nil, &auth, &code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/roku.txt", data, os.ModePerm)
}
