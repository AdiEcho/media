package main

import (
   "154.pages.dev/media/rtbf"
   "154.pages.dev/media/internal"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

func (f *flags) download() error {
   var (
      login rtbf.AuvioLogin
      err error
   )
   login.Raw, err = os.ReadFile(f.home + "/rtbf.txt")
   if err != nil {
      return err
   }
   err = login.Unmarshal()
   if err != nil {
      return err
   }
   token, err := login.Token()
   if err != nil {
      return err
   }
   auth, err := token.Auth()
   if err != nil {
      return err
   }
   address, err := func() (string, error) {
      u, err := url.Parse(f.address)
      if err != nil {
         return "", err
      }
      return u.Path, nil
   }()
   if err != nil {
      return err
   }
   var page rtbf.AuvioPage
   err = page.New(address)
   if err != nil {
      return err
   }
   title, err := auth.Entitlement(&page)
   if err != nil {
      return err
   }
   address, err = func() (string, error) {
      if v, ok := title.Dash(); ok {
         return v, nil
      }
      return "", errors.New("Entitlement.Dash")
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = &rtbf.Namer{page}
         f.s.Poster = title
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   var login rtbf.AuvioLogin
   err := login.New(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/rtbf.txt", login.Raw, os.ModePerm)
}
