package member

import (
   "fmt"
   "os"
   "reflect"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("cine_member"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*OperationUser).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", data, os.ModePerm)
}

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
