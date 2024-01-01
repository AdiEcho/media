package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/media/paramount"
   "154.pages.dev/stream"
   "fmt"
   "net/http"
   "os"
   "slices"
   "strings"
)

func (f flags) downloadable(token *paramount.App_Token) error {
   item, err := token.Item(f.content_ID)
   if err != nil {
      return err
   }
   ref, err := paramount.Downloadable(f.content_ID)
   if err != nil {
      return err
   }
   if f.s.Info {
      fmt.Println(ref)
      return nil
   }
   dst, err := os.Create(stream.Name(item) + ".mp4")
   if err != nil {
      return err
   }
   defer dst.Close()
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   src := log.New_Progress(1).Reader(res)
   if _, err := dst.ReadFrom(src); err != nil {
      return err
   }
   return nil
}

func (f flags) dash(token *paramount.App_Token) error {
   ref, err := paramount.DASH_CENC(f.content_ID)
   if err != nil {
      return err
   }
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f.s.Base = res.Request.URL
   reps, err := dash.Representations(res.Body)
   if err != nil {
      return err
   }
   if !f.s.Info {
      item, err := token.Item(f.content_ID)
      if err != nil {
         return err
      }
      f.s.Name, err = stream.Format_Film(item)
      if err != nil {
         return err
      }
      f.s.Poster, err = token.Session(f.content_ID)
      if err != nil {
         return err
      }
   }
   slices.SortFunc(reps, func(a, b *dash.Representation) int {
      return b.Bandwidth - a.Bandwidth
   })
   // video
   {
      reps := slices.Clone(reps)
      reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
         return !r.Video()
      })
      index := slices.IndexFunc(reps, func(r *dash.Representation) bool {
         if r.Height <= f.height {
            return r.Bandwidth <= f.bandwidth
         }
         return false
      })
      err := f.s.DASH_Sofia(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, func(r *dash.Representation) bool {
      return !r.Audio()
   })
   index := slices.IndexFunc(reps, func(r *dash.Representation) bool {
      if strings.HasPrefix(r.Codecs, f.codec) {
         if role, ok := r.Role(); ok {
            if role == f.role {
               return true
            }
         }
      }
      return false
   })
   return f.s.DASH_Sofia(reps, index)
}
