package nbc

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

const raw_pssh = "AAAAV3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADcIARIQBVLkSEJlSk6BsyYAS+R74BoLYnV5ZHJta2V5b3MiEAVS5EhCZUpOgbMmAEvke+AqAkhE"

func Test_License(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   pssh, err := base64.StdEncoding.DecodeString(raw_pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, nil, pssh)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(Core)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
