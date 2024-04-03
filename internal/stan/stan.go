package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/internal"
   "154.pages.dev/media/stan"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

type flags struct {
   media_id string
   h internal.HttpStream
   v log.Level
   web stan.WebAddress
   code bool
   auth bool
   secure bool
   home string
}

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

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.web, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.media_id, "i", "", "media ID")
   flag.BoolVar(&f.secure, "s", false, "secure URL")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.StringVar(&f.h.ClientId, "c", f.h.ClientId, "client ID")
   flag.StringVar(&f.h.PrivateKey, "p", f.h.PrivateKey, "private key")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   switch {
   case f.auth:
      err := f.write_auth()
      if err != nil {
         panic(err)
      }
   case f.code:
      err := f.write_code()
      if err != nil {
         panic(err)
      }
   case f.secure:
      err := f.write_secure()
      if err != nil {
         panic(err)
      }
   case f.web.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
func (f flags) download() error {
   var (
      secure stan.SecureUrl
      err error
   )
   secure.Data, err = os.ReadFile(f.web.String() + ".json")
   if err != nil {
      return err
   }
   secure.Unmarshal()
   // 1 VTT one
   for _, text := range secure.V.Text_Track_URLs {
      if text.ID == f.media_id {
         f.h.Name, err = f.web.Film()
         if err != nil {
            return err
         }
         return f.h.TimedText(text.URL)
      }
   }
   // 2 MPD one
   media, err := f.h.DashMedia(secure.V.URL)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.media_id {
         f.h.Name, err = f.web.Film()
         if err != nil {
            return err
         }
         var auth stan.Authenticate
         auth.Data, err = os.ReadFile(f.home + "/stan.json")
         if err != nil {
            return err
         }
         auth.Unmarshal()
         f.h.Poster = auth
         return f.h.DASH(medium)
      }
   }
   // 3 VTT all
   for _, text := range secure.V.Text_Track_URLs {
      fmt.Print(text, "\n\n")
   }
   // 4 MPD all
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}

func (f flags) write_auth() error {
   var (
      code stan.LinkCode
      err error
   )
   code.Data, err = os.ReadFile("link_code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/stan.json", auth.Data, 0666)
}

func (f flags) write_code() error {
   var code stan.LinkCode
   err := code.New()
   if err != nil {
      return err
   }
   os.WriteFile("link_code.json", code.Data, 0666)
   code.Unmarshal()
   fmt.Println(code)
   return nil
}

func (f flags) write_secure() error {
   var (
      auth stan.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/stan.json")
   if err != nil {
      return err
   }
   auth.Unmarshal()
   film, err := f.web.Film()
   if err != nil {
      return err
   }
   if err := auth.Viewing(film); err != nil {
      return err
   }
   secure, err := auth.URL(film)
   if err != nil {
      return err
   }
   return os.WriteFile(f.web.String() + ".json", secure.Data, 0666)
}
