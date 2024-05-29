package roku

import (
   "os"
   "testing"
)

func TestThree(t *testing.T) {
   text, err := os.ReadFile("2.json")
   if err != nil {
      t.Fatal(err)
   }
   var two two_response
   err = two.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   res, err := two.three()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
