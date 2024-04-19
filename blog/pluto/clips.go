package pluto

import (
   "encoding/json"
   "net/http"
   "strings"
)

func (o on_demand) clips() (episode_clips, error) {
   req, err := http.NewRequest("GET", "http://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      b.WriteString(o.ID)
      b.WriteString("/clips.json")
      return b.String()
   }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var clips episode_clips
   err = json.NewDecoder(res.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return clips, nil
}

func (b boot_start) video() (*on_demand, bool) {
   for _, video := range b.VOD {
      return &video, true
   }
   return nil, false
}

func (e episode_clips) clip() (*episode_clip, bool) {
   for _, clip := range e {
      return &clip, true
   }
   return nil, false
}

type episode_clips []episode_clip

type episode_clip struct {
   Sources []struct {
      File string
   }
}
