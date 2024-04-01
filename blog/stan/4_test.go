package stan

import (
   "fmt"
   "os"
   "testing"
)

func TestProgram(t *testing.T) {
   var (
      token web_token
      err error
   )
   token.data, err = os.ReadFile("2.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   session, err := token.session()
   if err != nil {
      t.Fatal(err)
   }
   program, err := session.program()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", program)
   fmt.Println(program.StanVideo())
}
