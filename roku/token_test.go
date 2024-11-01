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
   token, err := auth.Token(&code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", token.Raw, os.ModePerm)
}

func TestTokenRead(t *testing.T) {
   var err error
   // AccountToken
   var token AccountToken
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   token.Unmarshal()
   // AccountAuth
   var auth AccountAuth
   err = auth.New(&token)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
