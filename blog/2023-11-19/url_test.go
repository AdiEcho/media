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
      fmt.Println(path(test))
   }
}
