package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/hulu"
   "154.pages.dev/rosso"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   email string
   id hulu.ID
   level log.Level
   password string
   representation string
   s rosso.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Var(&f.id, "a", "address")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.password, "p", "", "password")
   flag.TextVar(&f.level, "v", f.level, "level")
   flag.StringVar(&f.s.Client_ID, "c", home+"client_id.bin", "client ID")
   flag.StringVar(&f.s.Private_Key, "k", home+"private_key.pem", "private key")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.id.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f flags) authenticate() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/hulu/token.json"
   auth, err := hulu.Living_Room(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(name, auth.Raw, 0666)
}
