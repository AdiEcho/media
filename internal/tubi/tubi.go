package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/tubi"
   "errors"
   "fmt"
   "net/http"
   "os"
)

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
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(rep, "\n\n")
      case rep.Id:
         f.s.Name = tubi.Namer{content}
         f.s.Poster = video
         return f.s.Download(rep)
      }
   }
   return nil
}
