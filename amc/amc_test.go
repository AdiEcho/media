package amc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

var path_tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
}

func Test_Path(t *testing.T) {
   for _, test := range path_tests {
      var u URL
      err := u.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(u)
   }
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
