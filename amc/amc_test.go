package amc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
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

func user(s string) (map[string]string, error) {
   b, err := os.ReadFile(s)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   json.Unmarshal(b, &m)
   return m, nil
}
