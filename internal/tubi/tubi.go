package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/tubi"
   "errors"
   "fmt"
   "net/http"
   "os"
   "sort"
)

func (f *flags) download() error {
   content := &tubi.VideoContent{}
   var err error
   content.Raw, err = os.ReadFile(f.name())
   if err != nil {
      return err
   }
   err = content.Unmarshal()
   if err != nil {
      return err
   }
   if content.Series() {
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("VideoContent.Get")
      }
   }
   video, ok := content.Video()
   if !ok {
      return errors.New("VideoContent.Video")
   }
   req, err := http.NewRequest("", video.Manifest.Url, nil)
   if err != nil {
      return err
   }
   reps, err := internal.Dash(req)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Name = tubi.Namer{content}
         f.s.Poster = video
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) name() string {
   return fmt.Sprint(f.tubi) + ".txt"
}

func (f *flags) write_content() error {
   content := &tubi.VideoContent{}
   err := content.New(f.tubi)
   if err != nil {
      return err
   }
   err = content.Unmarshal()
   if err != nil {
      return err
   }
   if content.Episode() {
      err := content.New(content.SeriesId)
      if err != nil {
         return err
      }
   }
   return os.WriteFile(f.name(), content.Raw, os.ModePerm)
}
