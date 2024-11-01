package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/rtbf"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
)

func (f *flags) authenticate() error {
   data, err := (*rtbf.AuvioLogin).New(nil, f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/rtbf.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/rtbf.txt")
   if err != nil {
      return err
   }
   var login rtbf.AuvioLogin
   err = login.Unmarshal(data)
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
   page, err := f.address.Page()
   if err != nil {
      return err
   }
   asset_id, ok := page.GetAssetId()
   if !ok {
      return errors.New("AuvioPage.GetAssetId")
   }
   title, err := auth.Entitlement(asset_id)
   if err != nil {
      return err
   }
   address, ok := title.Dash()
   if !ok {
      return errors.New("Entitlement.Dash")
   }
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Name = &rtbf.Namer{page}
         f.s.Poster = title
         return f.s.Download(rep)
      }
   }
   return nil
}
