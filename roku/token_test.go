package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestTokenWrite(t *testing.T) {
   // AccountAuth
   data, err := os.ReadFile("auth.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth AccountAuth
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   data, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   var code AccountCode
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountToken
   data, err = (*AccountToken).Marshal(nil, &auth, &code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
}

func TestTokenRead(t *testing.T) {
   // AccountToken
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AccountToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountAuth
   var auth AccountAuth
   data, err = auth.Marshal(&token)
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
