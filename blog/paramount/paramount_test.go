package paramount

import (
   "fmt"
   "testing"
)

// paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ
const content_id = "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ"

func TestMpd(t *testing.T) {
   address, err := DashCenc(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}
