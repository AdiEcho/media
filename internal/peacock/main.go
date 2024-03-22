package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "flag"
   "os"
   "path/filepath"
)

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.h.ClientId = f.home + "/widevine/client_id.bin"
   f.h.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   h internal.HttpStream
   home string
   id_session string
   media_id string
   peacock_id string
   v log.Level
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.peacock_id, "b", "", "Peacock ID")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.StringVar(&f.h.PrivateKey, "k", f.h.PrivateKey, "private key")
   flag.StringVar(&f.id_session, "s", "", "idsession")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.id_session != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.peacock_id != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
