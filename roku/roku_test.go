package roku

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "os"
   "testing"
   "time"
)

func TestPost(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.pssh != "" {
         pssh, err := base64.StdEncoding.DecodeString(test.pssh)
         if err != nil {
            t.Fatal(err)
         }
         mod, err := widevine.NewModule(private_key, client_id, nil, pssh)
         if err != nil {
            t.Fatal(err)
         }
         site, err := NewCrossSite()
         if err != nil {
            t.Fatal(err)
         }
         play, err := site.Playback(test.playback_id)
         if err != nil {
            t.Fatal(err)
         }
         key, err := mod.Key(play)
         if err != nil {
            t.Fatal(err)
         }
         if hex.EncodeToString(key) != test.key {
            t.Fatal(key)
         }
      }
   }
}

func TestPlayback(t *testing.T) {
   site, err := NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   for _, test := range tests {
      play, err := site.Playback(test.playback_id)
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(play)
      time.Sleep(time.Second)
   }
}

