package pluto

import (
   "net/http"
   "strings"
)

func (o on_demand) clips() (*http.Response, error) {
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
   return http.DefaultClient.Do(req)
}

func (b boot_start) video() (*on_demand, bool) {
   for _, video := range b.VOD {
      return &video, true
   }
   return nil, false
}
