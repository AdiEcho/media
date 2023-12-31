package youtube

import (
   "fmt"
   "testing"
)

func Test_YouTube(t *testing.T) {
   con, err := make_contents("2ZcDwdXEVyI")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(con)
}
