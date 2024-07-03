package paramount

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

var test = struct{
   content_id string
   key_id string
   url string
}{
   content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   key_id: "06c3b7eea1ce45779faee2abc8d01a55",
   url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
}

func TestWidevine(t *testing.T) {
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
   var pssh widevine.PSSH
   pssh.ContentId = []byte(test.content_id)
   pssh.KeyId, err = hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   var app AppToken
   app.com_cbs_app()
   session, err := app.Session(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(session, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestMpd(t *testing.T) {
   address, err := DashCenc(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}

func TestItem(t *testing.T) {
   var app AppToken
   err := app.com_cbs_ca()
   if err != nil {
      t.Fatal(err)
   }
   item, err := app.Item(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
