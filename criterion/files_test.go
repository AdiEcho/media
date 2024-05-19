package criterion

import (
   "os"
   "testing"
)

func TestFiles(t *testing.T) {
   var (
      token AuthToken
      err   error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   item, err := token.video(my_dinner)
   if err != nil {
      t.Fatal(err)
   }
   res, err := token.files(item)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
