package hulu

import (
   "154.pages.dev/http"
   "fmt"
   "testing"
)

func Test_Password(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   http.Verbose()
   auth, err := living_room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
