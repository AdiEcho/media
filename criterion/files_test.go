package criterion

import (
   "fmt"
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
   files, err := token.files(item)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", files)
}
