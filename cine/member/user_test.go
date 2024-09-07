package member

import (
   "fmt"
   "os"
   "reflect"
   "testing"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   ArticleAsset{},
   OperationArticle{},
   OperationPlay{},
   OperationUser{},
}

func TestAuthenticate(t *testing.T) {
   username := os.Getenv("cine_member_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("cine_member_password")
   var user OperationUser
   err := user.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", user.Raw, os.ModePerm)
}
