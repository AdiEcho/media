package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/youtube"
   "154.pages.dev/rosso"
   "fmt"
   "log/slog"
   "net/http"
   "os"
)

func (f flags) download() error {
   p, err := f.player()
   if err != nil {
      return err
   }
   slog.Info("playability", "status", p.PlayabilityStatus)
   forms := p.Streaming_Data.Adaptive_Formats
   if f.info {
      for i, form := range forms {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(form)
      }
   } else {
      var content youtube.Contents
      f.r.Web()
      content.Next(f.r)
      return encode(forms[index], rosso.Name(content))
   }
   return nil
}

func encode(f youtube.Format, name string) error {
   dst, err := func() (*os.File, error) {
      ext, err := f.Ext()
      if err != nil {
         return nil, err
      }
      return os.Create(name + ext)
   }()
   if err != nil {
      return err
   }
   defer dst.Close()
   log.TransportDebug()
   ranges := f.Ranges()
   var meter log.ProgressMeter
   meter.Set(len(ranges))
   for _, byte_range := range ranges {
      err := func() error {
         res, err := http.Get(f.URL + byte_range)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         _, err = dst.ReadFrom(meter.Reader(res))
         if err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) do_refresh() error {
   var code youtube.Device_Code
   code.Post()
   fmt.Println(code)
   fmt.Scanln()
   raw, err := code.Token()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return os.WriteFile(home+"/youtube.json", raw, 0666)
}

func (f flags) player() (*youtube.Player, error) {
   var token *youtube.Token
   switch f.request {
   case 0:
      f.r.Android()
   case 1:
      f.r.Android_Embed()
   case 2:
      f.r.Android_Check()
      home, err := os.UserHomeDir()
      if err != nil {
         return nil, err
      }
      raw, err := os.ReadFile(home + "/youtube.json")
      if err != nil {
         return nil, err
      }
      token = new(youtube.Token)
      token.Unmarshal(raw)
      if err := token.Refresh(); err != nil {
         return nil, err
      }
   }
   var play youtube.Player
   play.Post(f.r, token)
   return &play, nil
}
