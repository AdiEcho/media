package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/tubi"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "sort"
)

func (f *flags) write_content() error {
   var content tubi.VideoContent
   data, err := content.Marshal(f.tubi)
   if err != nil {
      return err
   }
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   if content.Episode() {
      data, err = content.Marshal(content.SeriesId)
      if err != nil {
         return err
      }
   }
   return os.WriteFile(f.name(), data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   content := &tubi.VideoContent{}
   err = content.Unmarshal(data)
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
   resp, err := http.Get(video.Manifest.Url)
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
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Name = tubi.Namer{content}
         f.s.Client = video
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) name() string {
   return fmt.Sprint(f.tubi) + ".txt"
}
