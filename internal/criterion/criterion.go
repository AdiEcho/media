package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/criterion"
   "154.pages.dev/media/internal"
   "flag"
   "fmt"
   "net/http"
   "os"
   "path/filepath"
)

type flags struct {
   email string
   s internal.Stream
   home string
   criterion criterion.ID
   representation string
   password string
   v log.Level
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.s.ClientId = f.home + "/widevine/client_id.bin"
   f.s.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.criterion, "a", "address")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.criterion.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f flags) authenticate() error {
   var auth criterion.Authenticate
   err := auth.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.json", auth.Data, 0666)
}

func (f flags) download() error {
   var (
      auth criterion.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/criterion.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.criterion)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
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

