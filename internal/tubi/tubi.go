package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/tubi"
   "fmt"
   "net/http"
   "os"
   "slices"
)

func (f *flags) download() error {
   text, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   var content tubi.Content
   err = content.Unmarshal(text)
   if err != nil {
      return err
   }
   video, ok := content.Video()
   if !ok {
      return errors.New("Content.Video")
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
         f.s.Name = tubi.Namer{&content}
         f.s.Poster = video
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) write_content() error {
   content := &tubi.Content{}
   err := content.New(f.tubi)
   if err != nil {
      return err
   }
   if content.Episode() {
      err := content.New(content.SeriesId)
      if err != nil {
         return err
      }
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return tubi.Content{}
      }
   }
   text, err := content.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.name(), text, os.ModePerm)
}

func (f *flags) name() string {
   return fmt.Sprint(f.tubi) + ".txt"
}
