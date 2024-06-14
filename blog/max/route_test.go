package max

import (
   "fmt"
   "testing"
)

var tests = []struct{
   class string
   path string
   url string
}{
   {
      class: "movie",
      path: "/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
      url: "play.max.com/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
   },
   {
      class: "episode",
      path: "/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68",
      url: "play.max.com/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68",
   },
}

func TestRoutes(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   routes, err := token.routes(tests[0].path)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", routes)
}
