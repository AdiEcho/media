package youtube

import (
   "154.pages.dev/http"
   "fmt"
   "testing"
   "time"
)

var ids = []string{
   "2ZcDwdXEVyI", // episode
   "HPkDFc8hq5c", // film
}

func Test_Watch(t *testing.T) {
   http.No_Location()
   http.Verbose()
   for _, id := range ids {
      c, err := make_contents(id)
      if err != nil {
         t.Fatal(err)
      }
      v, ok := c.show()
      fmt.Printf("%+v %v\n", v, ok)
      fmt.Println(c.title())
      time.Sleep(time.Second)
   }
}
