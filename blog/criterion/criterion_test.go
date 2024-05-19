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

func TestSlug(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   for _, a := range slug_a {
      for _, b := range slug_b {
         for _, c := range slug_c {
            for _, d := range slug_d {
               for _, e := range slug_e {
                  address := func() string {
                     var f strings.Builder
                     f.WriteString(a)
                     f.WriteString(b)
                     f.WriteString(c)
                     f.WriteString(d)
                     f.WriteString(e)
                     return f.String()
                  }()
                  status, err := token.do(address)
                  if err != nil {
                     t.Fatal(err)
                  }
                  fmt.Println(status, address)
                  time.Sleep(99 * time.Millisecond)
               }
            }
         }
      }
   }
}

func (a auth_token) do(address string) (string, error) {
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return "", err
   }
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
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
// criterionchannel.com/videos/my-dinner-with-andre
const video_id = 455774

func TestVideo(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   video, err := token.video(video_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
   name, err := encoding.Name(namer{video})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}

func TestToken(t *testing.T) {
   username := os.Getenv("criterion_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("criterion_password")
   var token auth_token
   err := token.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.data, 0666)
}
