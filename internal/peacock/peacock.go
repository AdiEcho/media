package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/peacock"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   dash_id string
   email string
   h internal.HttpStream
   peacock_id peacock.ID
   password string
   v log.Level
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.peacock_id, "a", "address")
   flag.StringVar(&f.h.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.dash_id, "d", "", "DASH ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.h.Private_Key, "k", home+"private_key.pem", "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.peacock_id.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth peacock.Authenticate
   auth.Raw, err = os.ReadFile(home + "/peacock.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   deep, err := auth.DeepLink(f.peacock_id)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   if f.dash_id != "" {
      detail, err := auth.Details(deep)
      if err != nil {
         return err
      }
      f.h.Name = detail
      f.h.Poster = play
   }
   media, err := f.h.DashMedia(play.Stream_URL)
   if err != nil {
      return err
   }
   return f.h.DASH(media, f.dash_id)
}

func (f flags) authenticate() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/peacock.json"
   auth, err := peacock.LivingRoom(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}
