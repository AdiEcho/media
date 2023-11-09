package hulu

import (
   "154.pages.dev/http"
   "fmt"
   "testing"
)

const id = "023c49bf-6a99-4c67-851c-4c9e7609cc1d"

func Test_Deeplink(t *testing.T) {
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
   link, err := auth.deeplink(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", link)
}
