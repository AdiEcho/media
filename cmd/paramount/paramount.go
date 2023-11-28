package main

import (
   "154.pages.dev/dash"
   "154.pages.dev/media/paramount"
   "154.pages.dev/stream"
   "fmt"
   "io"
   "net/http"
   "os"
   "slices"
   "strings"
   option "154.pages.dev/http"
)

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
      err := f.s.DASH_Get(reps, index)
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
   return f.s.DASH_Get(reps, index)
}

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
   name, err := stream.Format_Film(item)
   if err != nil {
      return err
   }
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(name + ".mp4")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := option.Progress_Length(res.ContentLength)
   if _, err := io.Copy(file, pro.Reader(res)); err != nil {
      return err
   }
   return nil
}

