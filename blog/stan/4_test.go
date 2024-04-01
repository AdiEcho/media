package stan

import (
   "fmt"
   "os"
   "testing"
)

// play.stan.com.au/programs/1768588
const program_id = 1768588

func TestProgram(t *testing.T) {
   var (
      token web_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   session, err := token.session()
   if err != nil {
      t.Fatal(err)
   }
   program, err := session.program(program_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", program)
   fmt.Printf("%+v\n", program.Media.DRM)
   fmt.Println(program.StanVideo())
}
