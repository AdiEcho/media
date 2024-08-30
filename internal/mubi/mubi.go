package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/mubi"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   var (
      secure mubi.SecureUrl
      err error
   )
   secure.Raw, err = os.ReadFile(f.address.String() + ".txt")
   if err != nil {
      return err
   }
   err = secure.Unmarshal()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", secure.Url, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, rep := range secure.TextTrackUrls {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         return f.timed_text(rep.Url)
      }
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         if _, ok := rep.Ext(); ok {
            fmt.Print(rep, "\n\n")
         }
      case rep.Id:
         film, err := f.address.Film()
         if err != nil {
            return err
         }
         f.s.Name = &mubi.Namer{film}
         var auth mubi.Authenticate
         auth.Raw, err = os.ReadFile(f.home + "/mubi.txt")
         if err != nil {
            return err
         }
         err = auth.Unmarshal()
         if err != nil {
            return err
         }
         f.s.Poster = &auth
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f flags) write_auth() error {
   var (
      code mubi.LinkCode
      err error
   )
   code.Raw, err = os.ReadFile("code.txt")
   if err != nil {
      return err
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/mubi.txt", auth.Raw, os.ModePerm)
}

func (f flags) write_code() error {
   var code mubi.LinkCode
   err := code.New()
   if err != nil {
      return err
   }
   os.WriteFile("code.txt", code.Raw, os.ModePerm)
   code.Unmarshal()
   fmt.Println(code)
   return nil
}

func (f flags) write_secure() error {
   var (
      auth mubi.Authenticate
      err error
   )
   auth.Raw, err = os.ReadFile(f.home + "/mubi.txt")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   film, err := f.address.Film()
   if err != nil {
      return err
   }
   err = auth.Viewing(film)
   if err != nil {
      return err
   }
   secure, err := auth.Url(film)
   if err != nil {
      return err
   }
   return os.WriteFile(f.address.String() + ".txt", secure.Raw, os.ModePerm)
}

func (f flags) timed_text(url string) error {
   resp, err := http.Get(url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   film, err := f.address.Film()
   if err != nil {
      return err
   }
   f.s.Name = &mubi.Namer{film}
   file, err := f.s.Create(".vtt")
   if err != nil {
      return err
   }
   defer file.Close()
   file.ReadFrom(resp.Body)
   return nil
}
