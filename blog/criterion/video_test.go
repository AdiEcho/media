package criterion

import (
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   res, err := token.video()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
