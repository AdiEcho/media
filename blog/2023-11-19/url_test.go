package url

import (
   "fmt"
   "testing"
)

var tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
}

func Test_URL(t *testing.T) {
   for _, test := range tests {
      p, err := path(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(p)
   }
}
