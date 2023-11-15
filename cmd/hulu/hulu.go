package main

import (
   "154.pages.dev/http"
   "154.pages.dev/media/nbc"
   "154.pages.dev/stream"
   "154.pages.dev/stream/dash"
   "flag"
   "net/http"
   "os"
   "path/filepath"
   "slices"
)

type flags struct {
   guid int64
   bandwidth int
   s stream.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home = filepath.ToSlash(home) + "/widevine/"
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.IntVar(&f.bandwidth, "bandwidth", 6_999_999, "maximum bandwidth")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Client_ID, "client", home+"client_id.bin", "client ID")
   flag.StringVar(&f.s.Private_Key, "key", home+"private_key.pem", "private key")
   flag.Parse()
   http.No_Location()
   http.Verbose()
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) download() error {
   meta, err := nbc.New_Metadata(f.guid)
   if err != nil {
      return err
   }
   reps, err := func() ([]dash.Representation, error) {
      o, err := meta.On_Demand()
      if err != nil {
         return nil, err
      }
      r, err := http.Get(o.Playback_URL)
      if err != nil {
         return nil, err
      }
      defer r.Body.Close()
      f.s.Base = r.Request.URL
      return dash.Representations(r.Body)
   }()
   if err != nil {
      return err
   }
   if !f.s.Info {
      f.s.Name, err = stream.Format_Episode(meta)
      if err != nil {
         return err
      }
      f.s.Poster = nbc.Core
   }
   // video
   {
      reps := slices.DeleteFunc(slices.Clone(reps), dash.Not(dash.Video))
      slices.SortFunc(reps, func(a, b dash.Representation) int {
         return b.Bandwidth - a.Bandwidth
      })
      index := slices.IndexFunc(reps, func(a dash.Representation) bool {
         return a.Bandwidth <= f.bandwidth
      })
      err := f.s.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, dash.Not(dash.Audio))
   return f.s.DASH_Get(reps, 0)
}
