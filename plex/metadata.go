package plex

import "net/url"

func (m metadata) dash(a anonymous) (*part, bool) {
   for _, media := range m.Media {
      if media.Protocol == "dash" {
         p := media.Part[0]
         p.Key = a.abs(p.Key, url.Values{})
         p.License = a.abs(p.License, url.Values{
            "x-plex-drm": {"widevine"},
         })
         return &p, true
      }
   }
   return nil, false
}

type metadata struct {
   Media []struct {
      Part []part
      Protocol string
   }
}

func (metadata) Show() string {
   return ""
}

func (metadata) Season() int {
   return 0
}

func (metadata) Episode() int {
   return 0
}

func (metadata) Title() string {
   return ""
}

func (metadata) Year() int {
   return 0
}
