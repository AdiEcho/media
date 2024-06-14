package max

import (
   "os"
   "testing"
)

// play.max.com/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68
const path = "/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68"

func TestRoute(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := token.route_android(path)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
