package criterion

import (
   "154.pages.dev/encoding"
   "fmt"
   "net/http"
   "os"
   "strings"
   "testing"
   "time"
)

func TestVideo(t *testing.T) {
   var (
      token AuthToken
      err   error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   res, err := token.video(my_dinner)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

func (a AuthToken) do(address string) (string, error) {
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return "", err
   }
   req.Header.Set("authorization", "Bearer "+a.v.AccessToken)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   return res.Status, nil
}

var slug_a = []string{
   "https://api.vhx.com",
   "https://api.vhx.tv",
}

var slug_b = []string{
   "",
   "/v2/sites/59054",
}

var slug_c = []string{
   "/collections/my-dinner-with-andre",
   "/videos/my-dinner-with-andre",
}

var slug_d = []string{
   "",
   "/items",
}

var slug_e = []string{
   "",
   "?site_id=59054",
}

func TestToken(t *testing.T) {
   username := os.Getenv("criterion_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("criterion_password")
   var token AuthToken
   err := token.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.data, 0666)
}

// criterionchannel.com/videos/my-dinner-with-andre
const my_dinner = "my-dinner-with-andre"

func TestItem(t *testing.T) {
   var (
      token AuthToken
      err   error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   err = token.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.item(my_dinner)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
   name, err := encoding.Name(item)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
