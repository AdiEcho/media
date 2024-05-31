package main

import (
   "154.pages.dev/media/internal"
   "154.pages.dev/media/tubi"
   "errors"
   "fmt"
   "net/http"
   "os"
)

func (f flags) download() error {
   text, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   var content tubi.Content
   err = content.Unmarshal(text)
   if err != nil {
      return err
   }
   video, err := content.Video()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("", video.Manifest.URL, nil)
   if err != nil {
      return err
   }
   media, err := internal.DASH(req)
   if err != nil {
      return err
   }
   for _, medium := range media {
      if medium.ID == f.representation {
         f.s.Poster = video
         f.s.Name = tubi.Namer{&content}
         return f.s.Download(medium)
      }
   }
   for i, medium := range media {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(medium)
   }
   return nil
}

func (f flags) name() string {
   return fmt.Sprint(f.tubi) + ".json"
}

func (f flags) write_content() error {
   content := new(tubi.Content)
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
         return errors.New("tubi.Content.Get")
      }
   }
   text, err := content.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.name(), text, 0666)
}
