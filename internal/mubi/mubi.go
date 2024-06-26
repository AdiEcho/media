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
   secure.Data, err = os.ReadFile(f.address.String() + ".json")
   if err != nil {
      return err
   }
   secure.Unmarshal()
   for _, text := range secure.V.TextTrackUrls {
      if text.ID == f.representation {
         film, err := f.address.Film()
         if err != nil {
            return err
         }
         f.s.Name = mubi.Namer{film}
         return f.s.TimedText(text.URL)
      }
   }
   req, err := http.NewRequest("", secure.V.URL, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.Id == f.representation {
         film, err := f.address.Film()
         if err != nil {
            return err
         }
         f.s.Name = mubi.Namer{film}
         var auth mubi.Authenticate
         auth.Data, err = os.ReadFile(f.home + "/mubi.json")
         if err != nil {
            return err
         }
         auth.Unmarshal()
         f.s.Poster = auth
         return f.s.Download(medium)
      }
   }
   // 3 VTT all
   for _, text := range secure.V.TextTrackUrls {
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
      code mubi.LinkCode
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
   return os.WriteFile(f.home + "/mubi.json", auth.Data, 0666)
}

func (f flags) write_code() error {
   var code mubi.LinkCode
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
      auth mubi.Authenticate
      err error
   )
   auth.Data, err = os.ReadFile(f.home + "/mubi.json")
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
   secure, err := auth.URL(film)
   if err != nil {
      return err
   }
   return os.WriteFile(f.address.String() + ".json", secure.Data, 0666)
}
