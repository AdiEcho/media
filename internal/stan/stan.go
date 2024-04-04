package main

import (
   "154.pages.dev/media/stan"
   "fmt"
   "os"
)

func (f flags) write_code() error {
   var code stan.ActivationCode
   err := code.New()
   if err != nil {
      return err
   }
   code.Unmarshal()
   fmt.Println(code)
   return os.WriteFile("code.json", code.Data, 0666)
}

func (f flags) write_token() error {
   var (
      code stan.ActivationCode
      err error
   )
   code.Data, err = os.ReadFile("code.json")
   if err != nil {
      return err
   }
   code.Unmarshal()
   token, err := code.Token()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/stan.json", token.Data, 0666)
}

func (f flags) download() error {
   var (
      token stan.WebToken
      err error
   )
   token.Data, err = os.ReadFile(f.home + "/stan.json")
   if err != nil {
      return err
   }
   token.Unmarshal()
   session, err := token.Session()
   if err != nil {
      return err
   }
   stream, err := session.Stream(f.program)
   if err != nil {
      return err
   }
   video, err := stream.StanVideo()
   if err != nil {
      return err
   }
   // 1 MPD one
   media, err := f.h.DashMedia(video.String())
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         var program stan.LegacyProgram
         err := program.New(f.program)
         if err != nil {
            return err
         }
         f.h.Name = program
         f.h.Poster = stream
         return f.h.DASH(medium)
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
