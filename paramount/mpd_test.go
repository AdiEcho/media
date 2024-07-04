package paramount

import (
   "fmt"
   "testing"
)

var tests = map[string]struct{
   content_id string
   key_id string
   url string
}{
   "fr": {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      key_id: "06c3b7eea1ce45779faee2abc8d01a55",
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
   "us": {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      key_id: "1fde0154d72a4f45912b34f0ce0777eb",
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
}

func TestMpdUs(t *testing.T) {
   address, err := MpegDash(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}

func TestMpdFr(t *testing.T) {
   address, err := MpegDash(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}
