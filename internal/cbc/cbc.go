package main

import (
   "154.pages.dev/encoding/hls"
   "154.pages.dev/media/cbc"
   "154.pages.dev/rosso"
   "os"
   "slices"
   "strings"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   profile, err := cbc.ReadProfile(home + "/cbc/profile.json")
   if err != nil {
      return nil, err
   }
   gem, err := cbc.NewCatalogGem(f.address)
   if err != nil {
      return nil, err
   }
   media, err := profile.Media(gem.Item())
   if err != nil {
      return nil, err
   }
   f.s.Name = rosso.Name(gem.StructuredMetadata)
   return f.s.HLS(media.URL)
   return f.s.HLS_Streams(master.Stream, index)
}

func (f flags) profile() error {
   login, err := cbc.NewToken(f.email, f.password)
   if err != nil {
      return err
   }
   profile, err := login.Profile()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return profile.WriteFile(home + "/cbc/profile.json")
}
