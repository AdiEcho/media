package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestTemplate(t *testing.T) {
   tmpl, err := new(template.Template).Parse(Template)
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Create("ignore.html")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   text, err := os.ReadFile("m3u8/desktop_master.m3u8")
   if err != nil {
      t.Fatal(err)
   }
   var master MasterPlaylist
   master.New(string(text))
   if err := tmpl.Execute(file, master); err != nil {
      t.Fatal(err)
   }
}

func TestId(t *testing.T) {
   for _, test := range id_tests {
      var req Request
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.VideoId)
   }
}

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

const image_test = "UpNXI3_ctAc"

func TestImage(t *testing.T) {
   for _, img := range Images {
      img.VideoId = image_test
      fmt.Println(img)
      res, err := http.Head(img.String())
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
