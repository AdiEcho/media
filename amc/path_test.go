package amc

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "fmt"
   "os"
   "testing"
   "time"
)

var path_tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
}

func Test_Path(t *testing.T) {
   for _, test := range path_tests {
      var p Path
      err := p.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(p)
   }
}

var tests = []struct {
   path string
   pssh string
} {
   { // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
      path: "/shows/orphan-black/episodes/season-1-instinct--1011152",
      pssh: "AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQuC5UBJ1cQS2w6wxWli1eSxoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=",
   },
   { // amcplus.com/movies/nocebo--1061554
      path: "/movies/nocebo--1061554",
   },
}

func Test_Content(t *testing.T) {
   var auth Auth_ID
   {
      s, err := os.UserHomeDir()
      if err != nil {
         t.Fatal(err)
      }
      b, err := os.ReadFile(s + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   http.No_Location()
   http.Verbose()
   for _, test := range tests {
      con, err := auth.Content(Path{test.path})
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Format_Film(vid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}
